package handlers

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
)

type SpreadSheetHandler struct {
	Inertia *inertia.Inertia
}

func NewSpreadSheetHandler(i *inertia.Inertia) *SpreadSheetHandler {
	return &SpreadSheetHandler{Inertia: i}
}

func (h *SpreadSheetHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "SpreadSheet/Index")

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
