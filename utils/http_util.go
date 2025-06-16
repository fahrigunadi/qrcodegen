package utils

import (
	"log"
	"net/http"
)

func HandleServerErr(w http.ResponseWriter, err error) {
	log.Printf("[HTTP ERROR] %s", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
