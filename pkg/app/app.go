package app

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Server struct {
	l *logrus.Logger
}

func New(l *logrus.Logger) Server {
	return Server{
		l: l,
	}
}

func (s *Server) HandleRomanNumeral() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query()["number"]
		if len(input) == 0 {

		}
		inputInt, err := strconv.Atoi(input[0])
	}
}
