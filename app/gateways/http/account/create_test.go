package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		name         string
		body         models.CreateRequest
		fields       account.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 201 successfully creating account",
			body: models.CreateRequest{Name: "Paulo", CPF: "55544433322", Secret: "125", BirthDate: "23/06/1994"},
			fields: account.UsecaseMock{CreateErr: nil, Account: vos.Account{
				ID:        "123",
				Name:      "Paulo",
				CPF:       "55544433322",
				CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
			}},

			wantStatus: http.StatusCreated,
			wantResponse: vos.CreateOutput{
				ID:        "123",
				Name:      "Paulo",
				CPF:       "55544433322",
				CreatedAt: "16/10/2022",
			},
		},
		{
			name:         "should return error 400 error when any required parameters are missing",
			body:         models.CreateRequest{Name: "", CPF: "33344433322", Secret: "12g", BirthDate: "12/04/2000"},
			fields:       account.UsecaseMock{CreateErr: nil, Account: vos.Account{}},
			wantStatus:   http.StatusBadRequest,
			wantResponse: models.ErrorResponse{Error: "create account request failed: a name is required to create the account"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.body)
			request, _ := http.NewRequest(http.MethodPost, "accounts", bytes.NewReader(requestBody))
			response := httptest.NewRecorder()

			h := Handler{
				UseCase: tt.fields,
			}

			h.Create(response, request)

			if tt.wantStatus != response.Code {
				t.Errorf("want: %d; got: %d", tt.wantStatus, response.Code)
			}

			wantBody, _ := json.Marshal(tt.wantResponse)
			if string(wantBody) != strings.TrimSpace(response.Body.String()) {
				t.Errorf("want: %s; got: %s", wantBody, response.Body.String())
			}
		})
	}
}
