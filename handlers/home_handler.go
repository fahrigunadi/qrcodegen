package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/fahrigunadi/qrcodegen/utils"
	inertia "github.com/romsar/gonertia"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type HomeHandler struct {
	Inertia *inertia.Inertia
}

func NewHomeHandler(i *inertia.Inertia) *HomeHandler {
	return &HomeHandler{Inertia: i}
}

func (h *HomeHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.Inertia.Render(w, r, "Home/Index")

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}

func (h *HomeHandler) Generate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data map[string]any
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			utils.HandleServerErr(w, err)
		}

		content, ok := data["content"].(string)
		if !ok {
			http.Error(w, "invalid content", http.StatusBadRequest)
			return
		}

		qr, err := qrcode.New(content)
		if err != nil {
			utils.HandleServerErr(w, err)
			return
		}

		var buf bytes.Buffer

		options := []standard.ImageOption{}

		if data["variant"] == "circle-shape" {
			options = append(options, standard.WithCircleShape())
		}

		writer := standard.NewWithWriter(utils.NopWriteCloser{Writer: &buf}, options...)
		if err := qr.Save(writer); err != nil {
			utils.HandleServerErr(w, err)
			return
		}

		base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
		dataURI := "data:image/png;base64," + base64Str

		resp := map[string]string{
			"base64": dataURI,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
}
