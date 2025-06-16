package main

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/handlers"
	"github.com/fahrigunadi/qrcodegen/internal/inertiasetup"
)

func main() {
	i := inertiasetup.Setup()

	mux := http.NewServeMux()

	homeHandler := handlers.NewHomeHandler(i)
	mux.Handle("GET /", i.Middleware(homeHandler.Index()))

	mux.Handle("GET /build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	http.ListenAndServe(":3000", mux)
}
