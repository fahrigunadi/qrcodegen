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
	port := getServerPort()

	i := inertiasetup.Setup()

	r := httprouter.New()

	r.ServeFiles("/build/*filepath", http.Dir("./public/build"))
	r.ServeFiles("/result/*filepath", http.Dir("./public/result"))

	r.GET("/", utils.Wrap(i.Middleware(handlers.NewHomeHandler(i).Index())))
	r.GET("/url", utils.Wrap(i.Middleware(handlers.NewUrlHandler(i).Index())))
	r.GET("/phone-number", utils.Wrap(i.Middleware(handlers.NewPhoneNumberHandler(i).Index())))
	r.GET("/email", utils.Wrap(i.Middleware(handlers.NewEmailHandler(i).Index())))
	r.GET("/spreadsheet", utils.Wrap(i.Middleware(handlers.NewSpreadSheetHandler(i).Index())))
	r.POST("/generate-qr", utils.Wrap(i.Middleware(handlers.NewGenerateHandler(i).Generate())))

	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Server running at http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getServerPort() string {
	port := flag.String("port", "3000", "Port to run the server on")
	short := flag.String("p", "3000", "Port to run the server on (shorthand)")
	flag.Parse()

	if *short != "3000" && *short != *port {
		return *short
	}
	return *port
}
