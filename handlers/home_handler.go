package handlers

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
)

type HomeHandler struct {
	Inertia *inertia.Inertia
}

func NewHomeHandler(i *inertia.Inertia) *HomeHandler {
	return &HomeHandler{Inertia: i}
}

func (h *HomeHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "Home/Index", inertia.Props{
			"text": "QR Code Generator",
		})
		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
