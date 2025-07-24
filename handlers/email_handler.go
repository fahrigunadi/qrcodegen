package handlers

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
)

type EmailHandler struct {
	Inertia *inertia.Inertia
}

func NewEmailHandler(i *inertia.Inertia) *EmailHandler {
	return &EmailHandler{Inertia: i}
}

func (h *EmailHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "Email/Index")

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
