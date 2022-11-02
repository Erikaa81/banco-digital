package transfer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/transfer"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/transfer/models"
)

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		name         string
		body         models.CreateRequest
		fields       transfer.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 201 successfully creating transfer",
			body: models.CreateRequest{AccountDestinationID: "2", Amount: 100},
			fields: transfer.UsecaseMock{
				CreateErr: nil,
				Transfer:  vos.Transfer{ID: "123", AccountOriginID: "1", AccountDestinationID: "2", Amount: 100, CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
			},
			wantStatus:   http.StatusCreated,
			wantResponse: vos.TransferOutput{ID: "123", AccountOriginID: "1", AccountDestinationID: "2", Amount: 100, CreatedAt: "16/10/2022"},
		},
		{
			name:         "should return error 400 error when any required parameters are missing",
			body:         models.CreateRequest{AccountDestinationID: "345", Amount: 0},
			fields:       transfer.UsecaseMock{CreateErr: nil, Transfer: vos.Transfer{}},
			wantStatus:   http.StatusBadRequest,
			wantResponse: models.ErrorResponse{Error: "create transfer request failed: amount invalid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.body)
			request, _ := http.NewRequest(http.MethodPost, "/accounts/account-id/transfers", bytes.NewReader(requestBody))
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
