package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type HandleError struct {
    Code int
    Message string
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (e HandleError) Error() string {
    return e.Message
}

func ToHttpHandler(handler HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Recover from panics
        defer func() {
            if rec := recover(); rec != nil {
                log.Printf("panic recovered: %v\n%s", rec, debug.Stack())
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()

        start := time.Now()

        err := handler(w, r)

        log.Printf("Method: %s, Path: %s, Duration: %v, Error: %v", r.Method, r.URL.Path, time.Since(start), err)

        if err != nil {
            switch e := err.(type) {
            case HandleError:
                SendError(w, r, e)
            default:
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }
    }
}

func SendError(w http.ResponseWriter, r *http.Request, err HandleError) {
    _ = r
    http.Error(w, err.Message, err.Code)
}

func SendJson(w http.ResponseWriter, r *http.Request, value any) error {
    _ = r
    var ok bool
    var err error
    var data []byte
    if data, ok = value.([]byte); !ok {
        data, err = json.Marshal(value)
        if err != nil {
            return err
        }
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return nil
}
