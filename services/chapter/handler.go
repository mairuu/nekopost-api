package chapter

import (
	"net/http"

	"github.com/mairuu/nekopost-api/api"
	"github.com/mairuu/nekopost-api/utils"
)

type Handler struct {
}

func NewHandler() Handler {
    return Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    http.HandleFunc("GET /api/chapters/latest", utils.ToHttpHandler(h.handleGetChapters))
    http.HandleFunc("GET /api/chapters/{chapter_id}/comments", utils.ToHttpHandler(h.handleGetChapterComments))
}

func (h *Handler) handleGetChapters(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    query := r.URL.Query()
    page := v.ValidateInt(query.Get("page"), "page")
    _type := v.ValidateProjectType(query.Get("type"))
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := api.GetChapters(page, _type)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}

func (h *Handler) handleGetChapterComments(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    cid := v.ValidateInt(r.PathValue("chapter_id"), "project id")
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := api.GetComments(cid)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}
