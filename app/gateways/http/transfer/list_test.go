package transfer

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/transfer"
	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/transfer/models"
)

func TestHandler_List(t *testing.T) {
	tests := []struct {
		name         string
		fields       transfer.UsecaseMock
		wantStatus   int
		wantResponse interface{}
	}{
		{
			name: "should return 200 successfully and transfers list",
			fields: transfer.UsecaseMock{
				GetByIDErr: nil,
				TransferList: []vos.Transfer{
					{ID: "1", AccountOriginID: "1", AccountDestinationID: "2", Amount: 100},
					{ID: "2", AccountOriginID: "1", AccountDestinationID: "3", Amount: 10},
					{ID: "3", AccountOriginID: "1", AccountDestinationID: "4", Amount: 150},
				},
			},
			wantStatus: http.StatusOK,
			wantResponse: models.GetIDResponse{
				List: []models.Transfer{
					{ID: "1", AccountOriginID: "1", AccountDestinationID: "2", Amount: 100},
					{ID: "2", AccountOriginID: "1", AccountDestinationID: "3", Amount: 10},
					{ID: "3", AccountOriginID: "1", AccountDestinationID: "4", Amount: 150},
				},
			},
		},
		{
			name: "should return 200 successfully and empty list",
			fields: transfer.UsecaseMock{
				TransferList: nil,
			},
			wantStatus: http.StatusOK,
			wantResponse: models.GetIDResponse{
				List: []models.Transfer{},
			},
		},
		{
			name: "should return error 500 when listing transfers fails",
			fields: transfer.UsecaseMock{
				GetByIDErr: errors.New("error when listing transfers"),
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: models.ErrorResponse{Error: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "accounts/{account-id}/transfers", bytes.NewReader(nil))
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
