package chapter

import (
	"net/http"

	"github.com/mairuu/nekopost-api/types"
	"github.com/mairuu/nekopost-api/utils"
)

type Handler struct {
    chapterApi types.ChapterApi;
    commentApi types.CommentApi;
}

func NewHandler(chapters types.ChapterApi) *Handler {
    return &Handler{
        chapterApi: chapters,
    }
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /api/chapters/latest", utils.ToHttpHandler(h.handleGetChapters))
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

    data, err := h.chapterApi.GetChapters(page, _type)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}

