package comment

import (
	"net/http"

	"github.com/mairuu/nekopost-api/types"
	"github.com/mairuu/nekopost-api/utils"
)

type Handler struct {
    commentApi types.CommentApi;
}

func NewHandler(comments types.CommentApi) *Handler {
    return &Handler{
        commentApi: comments,
    }
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /api/comments/{comment_id}", utils.ToHttpHandler(h.handleGetComment))
}

func (h *Handler) handleGetComment(w http.ResponseWriter, r *http.Request) error {
    v := utils.NewValidator()
    cid := v.ValidateInt(r.PathValue("comment_id"), "comment id")
    if err := v.Errors(); err != nil {
        return utils.HandleError{
            Code:    http.StatusBadRequest,
            Message: err.Error(),
        }
    }

    data, err := h.commentApi.GetComments(cid)
    if err != nil {
        return err
    }

    return utils.SendJson(w, r, data)
}
