package handlers

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
)

type PhoneNumberHandler struct {
	Inertia *inertia.Inertia
}

func NewPhoneNumberHandler(i *inertia.Inertia) *PhoneNumberHandler {
	return &PhoneNumberHandler{Inertia: i}
}

func (h *PhoneNumberHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "PhoneNumber/Index")

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
