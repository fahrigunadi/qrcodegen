package utils

import (
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func HandleServerErr(w http.ResponseWriter, err error) {
	log.Printf("[HTTP ERROR] %s", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func RedirectBack(w http.ResponseWriter, r *http.Request, fallback ...string) {
	referer := r.Header.Get("Referer")
	if referer == "" {
		if len(fallback) > 0 {
			referer = fallback[0]
		} else {
			referer = "/"
		}
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

func SetFlash(w http.ResponseWriter, name string, value string) {
	escaped := url.QueryEscape(value)

	cookie := &http.Cookie{
		Name:     name,
		Value:    escaped,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   300,
	}
	http.SetCookie(w, cookie)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string, dest *string) error {
	cookie, err := r.Cookie(name)
	if err != nil {
		return err
	}

	// Hapus cookie
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	unescaped, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return err
	}

	*dest = unescaped
	return nil
}

func RedirectBackWith(w http.ResponseWriter, r *http.Request, flashName string, flashData string) {
	SetFlash(w, flashName, flashData)

	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

func Wrap(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}
