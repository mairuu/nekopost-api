package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandleError struct {
    Code int
    Message string
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (e HandleError) Error() string {
    return e.Message
}

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := f(w, r)
    if err == nil {
        return
    }
    herr, ok := err.(HandleError)
    if !ok {
        herr = HandleError{
            Code: http.StatusInternalServerError,
            Message: http.StatusText(http.StatusInternalServerError),
        }
    }
    SendError(w, r, herr)
}

func ToHttpHandler(handler HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := handler(w, r)
        if err == nil {
            return
        }
        herr, ok := err.(HandleError)
        if !ok {
            herr = HandleError{
                Code: http.StatusInternalServerError,
                Message: http.StatusText(http.StatusInternalServerError),
            }
        }
        SendError(w, r, herr)
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
