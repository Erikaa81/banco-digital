package transfer

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/transfer"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/transfer/models"
)

func TestHandler_GetByIDTransfer(t *testing.T) {
	tests := []struct {
		name         string
		body         models.GetIDTransferRequest
		fields       transfer.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 200 successfully transfer",
			body: models.GetIDTransferRequest{ID: "1"},
			fields: transfer.UsecaseMock{
				CreateErr: nil,
				Transfer: vos.Transfer{
					ID:                   "1",
					AccountOriginID:      "223",
					AccountDestinationID: "112",
					Amount:               1000,
					CreatedAt:            time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
				},
			},
			wantStatus: http.StatusOK,
			wantResponse: models.GetIDTransferResponse{
				ID:                   "1",
				AccountOriginID:      "223",
				AccountDestinationID: "112",
				Amount:               1000,
				CreatedAt:            "16/10/2022",
			},
		},
		{
			name: "should return error 404 when looking for id at repository",
			fields: transfer.UsecaseMock{
				GetByIDTransferErr: ErrIDNotFound,
			},
			wantStatus:   http.StatusNotFound,
			wantResponse: models.ErrorResponse{Error: "id not found"},
		},
		{
			name: "should return error 500 search by id failed",
			fields: transfer.UsecaseMock{
				GetByIDTransferErr: errors.New("unexpected error"),
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: models.ErrorResponse{Error: "search by id failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.body)
			request, _ := http.NewRequest(http.MethodGet, "accounts/{account-id}/transfers/{transfer-id}", bytes.NewReader(requestBody))
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
