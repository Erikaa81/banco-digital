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

func TestHandler_List(t *testing.T) {
	tests := []struct {
		name         string
		fields       account.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 200 successfully and account list",
			fields: account.UsecaseMock{
				ListErr: nil,
				AccountsList: []vos.Account{
					{ID: "1", Name: "Camila", CPF: "22233344455", BirthDate: time.Date(2003, 10, 2, 0, 0, 0, 0, time.Local)},
					{ID: "2", Name: "Maria", CPF: "55533344455", BirthDate: time.Date(2000, 19, 7, 0, 0, 0, 0, time.Local)},
					{ID: "3", Name: "Paula", CPF: "44455566647", BirthDate: time.Date(1980, 26, 9, 0, 0, 0, 0, time.Local)},
				},
			},
			wantStatus: http.StatusOK,
			wantResponse: models.ListResponse{
				List: []models.Account{
					{ID: "1", Name: "Camila", CPF: "22233344455", BirthDate: time.Date(2003, 10, 2, 0, 0, 0, 0, time.Local)},
					{ID: "2", Name: "Maria", CPF: "55533344455", BirthDate: time.Date(2000, 19, 7, 0, 0, 0, 0, time.Local)},
					{ID: "3", Name: "Paula", CPF: "44455566647", BirthDate: time.Date(1980, 26, 9, 0, 0, 0, 0, time.Local)},
				},
			},
		},
		{
			name: "should return 200 successfully and empty list",
			fields: account.UsecaseMock{
				AccountsList: nil,
			},
			wantStatus: http.StatusOK,
			wantResponse: models.ListResponse{
				List: []models.Account{},
			},
		},
		{
			name: "should return error 500 when listing accounts fails",
			fields: account.UsecaseMock{
				ListErr:      errors.New("error when listing accounts"),
			},
			wantStatus: http.StatusInternalServerError,
			wantResponse: models.ErrorResponse{Error: "error when listing accounts"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "accounts", bytes.NewReader(nil))
			response := httptest.NewRecorder()

			h := Handler{
				UseCase: tt.fields,
			}

			h.List(response, request)
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
