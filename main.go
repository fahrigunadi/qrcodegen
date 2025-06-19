package main

import (
	"net/http"

	"github.com/fahrigunadi/qrcodegen/handlers"
	"github.com/fahrigunadi/qrcodegen/internal/inertiasetup"
	"github.com/fahrigunadi/qrcodegen/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	i := inertiasetup.Setup()

	r := httprouter.New()

	homeHandler := handlers.NewHomeHandler(i)
	urlHandler := handlers.NewUrlHandler(i)

	r.ServeFiles("/build/*filepath", http.Dir("./public/build"))
	r.ServeFiles("/result/*filepath", http.Dir("./public/result"))

	r.GET("/", utils.Wrap(i.Middleware(homeHandler.Index())))
	r.GET("/url", utils.Wrap(i.Middleware(urlHandler.Index())))
	r.POST("/generate-qr", utils.Wrap(i.Middleware(homeHandler.Generate())))

	http.ListenAndServe(":3000", r)
}
