package models

import "testing"

func TestCreateRequest_Validate(t *testing.T) {
	type fields struct {
		AccountOriginID      string
		AccountDestinationID string
		Amount               int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name:    "should return error because account destination is mandatory",
			fields:  fields{AccountOriginID: "1234", AccountDestinationID: "", Amount: 100},
			wantErr: ErrMandatoryAccountDestination,
		},
		{
			name:    "should return error because amount invalid",
			fields:  fields{AccountOriginID: "1234", AccountDestinationID: "213", Amount: -10},
			wantErr: ErrAmountInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &CreateRequest{
				AccountDestinationID: tt.fields.AccountDestinationID,
				Amount:               tt.fields.Amount,
			}
			if err := tr.Validate(); err != tt.wantErr {
				t.Errorf("CreateRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
