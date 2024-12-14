package project

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mairuu/nekopost-api/api"
	"github.com/mairuu/nekopost-api/utils"
)

type Handler struct {
}

func NewHandler() Handler {
    return Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /api/projects/random", utils.ToHttpHandler(h.handleGetRandomProjects))
    mux.HandleFunc("GET /api/projects/popular", utils.ToHttpHandler(h.handleGetPopularProjects))
    mux.HandleFunc("GET /api/projects", utils.ToHttpHandler(h.handleGetProjects))
    mux.HandleFunc("GET /api/projects/{project_id}", utils.ToHttpHandler(h.handleGetProject))
    mux.HandleFunc("GET /api/projects/{project_id}/cover.jpg", utils.ToHttpHandler(h.handleGetProjectCover))
    mux.HandleFunc("GET /api/projects/{project_id}/chapters/{chapter_id}", utils.ToHttpHandler(h.handleGetProjectChapter))
}

func (h *Handler) handleGetRandomProjects(w http.ResponseWriter, r *http.Request) error {
    data, err := api.GetRandomProjects()
    if err != nil {
        return err
    }
    return utils.SendJson(w, r, data)
}

func (h *Handler) handleGetPopularProjects(w http.ResponseWriter, r *http.Request) error {
    data, err := api.GetPopularProjects()
    if err != nil {
        return err
    }
    return utils.SendJson(w, r, data)
}

func (h *Handler) handleGetProjects(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    query := r.URL.Query()
    order := "project"
    page := v.ValidateInt(query.Get("page"), "page")
    types := v.ValidateProjectTypes(query.Get("types"))
    genres := v.ValidateGenres(query.Get("genres"))
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := api.GetProjects(page, order, types, genres)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}

func (h *Handler) handleGetProject(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := api.GetProject(pid)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}

func (h *Handler) handleGetProjectCover(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
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

func (h *Handler) handleGetProjectChapter(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    pid := v.ValidateInt(r.PathValue("project_id"), "project id")
    cid := v.ValidateInt(r.PathValue("chapter_id"), "chapter id")
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := api.GetChapter(pid, cid)
    if err != nil {
        return err
    }
    return utils.SendJson(w, r, data)
}
