package main

import (
	"log"
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

	r.ServeFiles("/build/*filepath", http.Dir("./public/build"))
	r.ServeFiles("/result/*filepath", http.Dir("./public/result"))

	r.GET("/", utils.Wrap(i.Middleware(homeHandler.Index())))
	r.POST("/generate-qr", utils.Wrap(i.Middleware(homeHandler.Generate())))

	log.Fatal(http.ListenAndServe(":3001", r))
}
