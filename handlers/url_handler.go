package handlers

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
)

type UrlHandler struct {
	Inertia *inertia.Inertia
}

func NewUrlHandler(i *inertia.Inertia) *UrlHandler {
	return &UrlHandler{Inertia: i}
}

func (h *UrlHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "Url/Index")

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
