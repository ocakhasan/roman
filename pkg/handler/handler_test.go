package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ocakhasan/roman/pkg/handler/structure"

	handlerError "github.com/ocakhasan/roman/pkg/handler/error"

	"github.com/stretchr/testify/assert"

	"github.com/ocakhasan/roman/pkg/app"
	"github.com/sirupsen/logrus"
)

func newTestLogger() *logrus.Logger {
	return logrus.New()
}

func newTestServer() *httptest.Server {
	l := newTestLogger()

	appTy := app.New(l)

	handlers := CreateHandler(appTy)

	return httptest.NewServer(handlers)
}

func TestHandler_RomanNumeralConversion(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()
	client := ts.Client()

	cases := []struct {
		name             string
		httpMethod       string
		query            string
		statusCode       int
		expectedResponse interface{}
		errHappened      bool
	}{
		{
			name:       "empty query input",
			httpMethod: http.MethodPost,
			query:      "",
			statusCode: http.StatusBadRequest,
			expectedResponse: handlerError.ErrorResponse{
				Err:        handlerError.ErrEmptyInput.Error(),
				StatusCode: http.StatusBadRequest,
			},
			errHappened: true,
		},
		{
			name:       "invalid query input",
			httpMethod: http.MethodPost,
			query:      "asd",
			statusCode: http.StatusBadRequest,
			expectedResponse: handlerError.ErrorResponse{
				Err:        handlerError.ErrInvalidInput.Error(),
				StatusCode: http.StatusBadRequest,
			},
			errHappened: true,
		},
		{
			name:       "valid input",
			httpMethod: http.MethodPost,
			query:      "12",
			statusCode: http.StatusOK,
			expectedResponse: structure.RomanResponse{
				Input:  "12",
				Output: "XII",
			},
			errHappened: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/romannumeral", ts.URL)
			if tt.query != "" {
				url = fmt.Sprintf("%s?query=%s", url, tt.query)
			}

			request, err := http.NewRequest(tt.httpMethod, url, nil)
			if err != nil {
				t.Errorf("unexpected error %v\n", err)
			}

			request.Header.Set("Referer", "http://localhost")

			response, err := client.Do(request)
			if err != nil {
				t.Errorf("unexpected response error %v\n", err)
			}

			if response.StatusCode != tt.statusCode {
				t.Errorf("expected statusCode :%v, got: %v\n", tt.statusCode, response.StatusCode)
			}

			if tt.errHappened {
				var resp handlerError.ErrorResponse
				if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Errorf("error while decoding the error response body: %v", err)
				}
				assert.Equal(t, resp, tt.expectedResponse)
			} else {
				var resp structure.RomanResponse
				if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Errorf("error while decoding the vehicle response body: %v", err)
				}
				assert.Equal(t, resp, tt.expectedResponse)
			}
		})
	}
}

func TestHandler_MinMaxRangeConversions(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()
	client := ts.Client()

	cases := []struct {
		name             string
		httpMethod       string
		minQuery         string
		maxQuery         string
		statusCode       int
		expectedResponse interface{}
		errHappened      bool
	}{
		{
			name:       "empty min query input",
			httpMethod: http.MethodPost,
			minQuery:   "",
			maxQuery:   "12",
			statusCode: http.StatusBadRequest,
			expectedResponse: handlerError.ErrorResponse{
				Err:        handlerError.ErrEmptyInput.Error(),
				StatusCode: http.StatusBadRequest,
			},
			errHappened: true,
		},
		{
			name:       "invalid min or max query input",
			httpMethod: http.MethodPost,
			minQuery:   "1",
			maxQuery:   "12asdas",
			statusCode: http.StatusBadRequest,
			expectedResponse: handlerError.ErrorResponse{
				Err:        handlerError.ErrInvalidInput.Error(),
				StatusCode: http.StatusBadRequest,
			},
			errHappened: true,
		},
		{
			name:       "min is bigger than max",
			httpMethod: http.MethodPost,
			minQuery:   "12",
			maxQuery:   "2",
			statusCode: http.StatusBadRequest,
			expectedResponse: handlerError.ErrorResponse{
				Err:        handlerError.ErrMinBiggerThanMax.Error(),
				StatusCode: http.StatusBadRequest,
			},
			errHappened: true,
		},
		{
			name:       "valid min max query",
			httpMethod: http.MethodPost,
			minQuery:   "1",
			maxQuery:   "5",
			statusCode: http.StatusOK,
			expectedResponse: structure.RangeResponse{Conversions: []structure.RomanResponse{
				{
					Input:  "1",
					Output: "I",
				},
				{
					Input:  "2",
					Output: "II",
				},
				{
					Input:  "3",
					Output: "III",
				},
				{
					Input:  "4",
					Output: "IV",
				},
				{
					Input:  "5",
					Output: "V",
				},
			}},
			errHappened: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/romannumeral", ts.URL)
			if tt.minQuery != "" {
				url = fmt.Sprintf("%s?min=%s&", url, tt.minQuery)
			}

			if tt.maxQuery != "" {
				if tt.minQuery == "" {
					url = fmt.Sprintf("%s?max=%s", url, tt.maxQuery)
				} else {
					url = fmt.Sprintf("%smax=%s", url, tt.maxQuery)
				}
			}

			request, err := http.NewRequest(tt.httpMethod, url, nil)
			if err != nil {
				t.Errorf("unexpected error %v\n", err)
			}

			request.Header.Set("Referer", "http://localhost")

			response, err := client.Do(request)
			if err != nil {
				t.Errorf("unexpected response error %v\n", err)
			}

			if response.StatusCode != tt.statusCode {
				t.Errorf("expected statusCode :%v, got: %v\n", tt.statusCode, response.StatusCode)
			}

			if tt.errHappened {
				var resp handlerError.ErrorResponse
				if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Errorf("error while decoding the error response body: %v", err)
				}
				assert.Equal(t, resp, tt.expectedResponse)
			} else {
				var resp structure.RangeResponse
				if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
					t.Errorf("error while decoding the vehicle response body: %v", err)
				}
				assert.Equal(t, resp, tt.expectedResponse)
			}
		})
	}
}
