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

		input := r.URL.Query()["query"]
		if len(input) == 0 { // If query is not provided, it is assumed that it is a min-max query
			minQuery := r.URL.Query()["min"]
			maxQuery := r.URL.Query()["max"]

			if len(minQuery) == 0 || len(maxQuery) == 0 {
				s.l.Error("min or max value is not provided")
				handlerError.WriteError(w, handlerError.ErrInvalidInput)
				return
			}

			minInt, maxInt, err := validateMinMax(minQuery[0], maxQuery[0])
			if err != nil {
				s.l.Errorf("error while validating min max %v", err)
				handlerError.WriteError(w, err)
				return
			}

			output := roman.NumeralRange(minInt, maxInt)
			response := structure.RangeResponse{
				Conversions: output,
			}

			if err := json.NewEncoder(w).Encode(&response); err != nil {
				s.l.Errorf("error while encoding the response %v", err)
				handlerError.WriteError(w, err)
				return
			}

		} else {
			inputInt, err := strconv.Atoi(input[0])
			if err != nil {
				s.l.Errorf("Invalid input format %v", input[0])
				handlerError.WriteError(w, handlerError.ErrInvalidInput)
				return
			}

			output := roman.ConvertIntegerToRoman(inputInt)

			response := structure.RomanResponse{
				Input:  input[0],
				Output: output,
			}

			if err := json.NewEncoder(w).Encode(&response); err != nil {
				s.l.Errorf("error while encoding the response %v", err)
				handlerError.WriteError(w, err)
				return
			}
		}

	}
}

func validateMinMax(min, max string) (int, int, error) {
	minInt, err := strconv.Atoi(min)
	if err != nil {
		return 0, 0, handlerError.ErrInvalidInput
	}

	maxInt, err := strconv.Atoi(max)
	if err != nil {
		return 0, 0, handlerError.ErrInvalidInput
	}

	if minInt > maxInt {
		return 0, 0, handlerError.ErrMinBiggerThanMax
	}

	return minInt, maxInt, nil
}
