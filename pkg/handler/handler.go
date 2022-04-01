package handler

import (
	"github.com/ocakhasan/roman/pkg/app"
	"net/http"
)

func CreateHandler(app app.Server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/romannumeral", app.HandleRomanNumeral())

	return mux
}
