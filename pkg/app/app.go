package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ocakhasan/roman/pkg/handler/structure"

	"github.com/ocakhasan/roman/internal/roman"

	handlerError "github.com/ocakhasan/roman/pkg/handler/error"

	"github.com/sirupsen/logrus"
)

const (
	contentType = "application/json"
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
		w.Header().Set("content-type", contentType)
		input := r.URL.Query()["number"]
		if len(input) == 0 {
			s.l.Error("Empty query name")
			handlerError.WriteError(w, handlerError.ErrEmptyInput)
			return
		}
		inputInt, err := strconv.Atoi(input[0])
		if err != nil {
			s.l.Errorf("Invalid input format %v", input[0])
			handlerError.WriteError(w, handlerError.ErrInvalidInput)
			return
		}

		output, err := roman.ConvertIntegerToRoman(inputInt)
		if err != nil {
			s.l.Error("error while converting to roman %v", err)
			handlerError.WriteError(w, err)
			return
		}

		response := structure.RomanResponse{
			Input:  input[0],
			Output: output,
		}

		json.NewEncoder(w).Encode(&response)

	}
}
