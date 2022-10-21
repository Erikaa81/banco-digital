package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

func TestHandler_GetByID(t *testing.T) {
	tests := []struct {
		name         string
		body         models.GetIDRequest
		fields       account.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 200 successfully account",
			body: models.GetIDRequest{ID: "123"},
			fields: account.UsecaseMock{
				CreateErr: nil,
				Account: vos.Account{
					ID:        "123",
					Name:      "Paulo",
					CPF:       "55544433322",
					CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
				},
			},
			wantStatus: http.StatusOK,
			wantResponse: models.GetIDResponse{
				ID:        "123",
				Name:      "Paulo",
				CPF:       "55544433322",
				CreatedAt: "16/10/2022",
			},
		},
		{
			name: "should return error 404 when looking for id at repository",
			fields: account.UsecaseMock{
				GetByIDErr: ErrIDNotFound,
			},
			wantStatus:   http.StatusNotFound,
			wantResponse: models.ErrorResponse{Error: "id not found"},
		},
		{
			name: "should return error 500 search by id failed",
			fields: account.UsecaseMock{
				GetByIDErr: errors.New("unexpected error"),
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: models.ErrorResponse{Error: "search by id failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.body)
			request, _ := http.NewRequest(http.MethodGet, "accounts/{id}", bytes.NewReader(requestBody))
			response := httptest.NewRecorder()

			h := Handler{
				UseCase: tt.fields,
			}

			h.GetByID(response, request)
			if tt.wantStatus != response.Code {
				t.Errorf("want: %d; got: %d", tt.wantStatus, response.Code)
			}

			wantBody, _ := json.Marshal(tt.wantResponse)
			if string(wantBody) != strings.TrimSpace(response.Body.String()) {
				t.Errorf("want: %s; got: %s", wantBody, response.Body)
			}
		})
	}
}
