package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

func TestHandler_GetBalance(t *testing.T) {
	tests := []struct {
		name         string
		body         models.GetBalanceRequest
		fields       account.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 200 successfully balance",
			body: models.GetBalanceRequest{ID: "123"},
			fields: account.UsecaseMock{
				CreateErr: nil,
				Account: vos.Account{
					ID:      "123",
					Name:    "Paulo",
					Balance: 100,
				},
			},
			wantStatus: http.StatusOK,
			wantResponse: models.GetBalanceResponse{
				Balance: 100,
			},
		},
		{
			name: "should return error 404 when looking for id at repository",
			fields: account.UsecaseMock{
				GetBalanceErr: ErrIDNotFound,
			},
			wantStatus:   http.StatusNotFound,
			wantResponse: models.ErrorResponse{Error: "id not found"},
		},
		{
			name: "should return error 500 search balance failed",
			fields: account.UsecaseMock{
				GetBalanceErr: errors.New("unexpected error"),
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: models.ErrorResponse{Error: "search balance failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.body)
			request, _ := http.NewRequest(http.MethodGet, "accounts/{id}/balance", bytes.NewReader(requestBody))
			response := httptest.NewRecorder()

			h := Handler{
				UseCase: tt.fields,
			}

			h.GetBalance(response, request)
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
