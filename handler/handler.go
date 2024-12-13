package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mairuu/nekopost-api/api"
	"github.com/mairuu/nekopost-api/validator"
)

type HandleError struct {
    code int
    message string
}

func (e HandleError) Error() string {
    return e.message
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

var HandleLatestChapters = toHttpHandler(handleChapters)
var HandleChapterComments = toHttpHandler(handleChapterComments)

var HandleRandomProjects = toHttpHandler(handleRandomProjects)
var HandlePopularProjects = toHttpHandler(handlePopularProjects)
var HandleProjects = toHttpHandler(handleProjects)
var HandleProjectCover = toHttpHandler(handleProjectCover)
var HandleProject = toHttpHandler(handleProject)
var HandleProjectChapter = toHttpHandler(handleProjectChapter)

func toHttpHandler(handler HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := handler(w, r)
        if err == nil {
            return
        }
        herr, ok := err.(HandleError)
        if !ok {
            herr = HandleError{
                code: http.StatusInternalServerError,
                message: http.StatusText(http.StatusInternalServerError),
            }
        }
        sendError(w, r, herr)
    }
}

func sendError(w http.ResponseWriter, r *http.Request, err HandleError) {
    _ = r
    w.WriteHeader(err.code)
    w.Write([]byte(err.message))
}

func sendJson(w http.ResponseWriter, r *http.Request, value any) error {
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


func handleChapters(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    query := r.URL.Query()
    page := v.ValidateInt(query.Get("page"), "page")
    _type := v.ValidateProjectType(query.Get("type"))
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    data, err := api.GetChapters(page, _type)
    if err != nil {
        return err
    }

    return sendJson(w, r, data)
}

func handleChapterComments(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    cid := v.ValidateInt(r.PathValue("chapter_id"), "project id")
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    data, err := api.GetComments(cid)
    if err != nil {
        return err
    }

    return sendJson(w, r, data)
}

func handleRandomProjects(w http.ResponseWriter, r *http.Request) error {
    data, err := api.GetRandomProjects()
    if err != nil {
        return err
    }
    return sendJson(w, r, data)
}

func handlePopularProjects(w http.ResponseWriter, r *http.Request) error {
    data, err := api.GetPopularProjects()
    if err != nil {
        return err
    }
    return sendJson(w, r, data)
}

func handleProjects(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    query := r.URL.Query()
    order := "project"
    page := v.ValidateInt(query.Get("page"), "page")
    types := v.ValidateProjectTypes(query.Get("types"))
    genres := v.ValidateGenres(query.Get("genres"))
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    data, err := api.GetProjects(page, order, types, genres)
    if err != nil {
        return err
    }

    return sendJson(w, r, data)
}

func handleProject(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    data, err := api.GetProject(pid)
    if err != nil {
        return err
    }

    return sendJson(w, r, data)
}

func handleProjectCover(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    coverUrl := fmt.Sprintf("https://www.osemocphoto.com/collectManga/%d/%d_cover.jpg", pid, pid)
    req, err := http.NewRequest(http.MethodGet, coverUrl, nil)
    if err != nil {
        return err 
    }
    req.Header.Set("Referer", "https://www.nekopost.net/")
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    res, err := client.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    // forward headers
    w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
    w.Header().Set("Cache-Control", "public, max-age=604800") // 7 days

    io.Copy(w, res.Body)
    return nil
}

func handleProjectChapter(w http.ResponseWriter, r *http.Request) error {
    v := validator.New()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    cid := v.ValidateInt(r.PathValue("chapter_id"), "chapter id")
    if err := v.Errors(); err != nil {
        return HandleError{
            code:    http.StatusBadRequest,
            message: err.Error(),
        }
    }

    data, err := api.GetChapter(pid, cid)
    if err != nil {
        return err
    }
    return sendJson(w, r, data)
}
