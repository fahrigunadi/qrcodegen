package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fahrigunadi/qrcodegen/handlers"
	"github.com/fahrigunadi/qrcodegen/internal/inertiasetup"
	"github.com/fahrigunadi/qrcodegen/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := flag.String("port", "3000", "Port to run the server on")
	p := flag.String("p", "3000", "Port to run the server on (shorthand)")

	flag.Parse()

	selectedPort := *port
	if *p != "3000" && *p != *port {
		selectedPort = *p
	}

	i := inertiasetup.Setup()

	r := httprouter.New()

	homeHandler := handlers.NewHomeHandler(i)
	urlHandler := handlers.NewUrlHandler(i)
	phoneNumberHandler := handlers.NewPhoneNumberHandler(i)

	r.ServeFiles("/build/*filepath", http.Dir("./public/build"))
	r.ServeFiles("/result/*filepath", http.Dir("./public/result"))

	r.GET("/", utils.Wrap(i.Middleware(homeHandler.Index())))
	r.GET("/url", utils.Wrap(i.Middleware(urlHandler.Index())))
	r.GET("/phone-number", utils.Wrap(i.Middleware(phoneNumberHandler.Index())))
	r.POST("/generate-qr", utils.Wrap(i.Middleware(homeHandler.Generate())))

	addr := fmt.Sprintf(":%s", selectedPort)
	fmt.Println("Server running at http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
