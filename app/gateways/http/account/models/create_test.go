package models

import (
	"testing"
)

func TestCreateRequest_Validate(t *testing.T) {
	type fields struct {
		Name      string
		CPF       string
		Secret    string
		BirthDate string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{

		{
			name:    "should return error because name is mandatory",
			fields:  fields{Name: "", CPF: "33222211122", Secret: "76t"},
			wantErr: ErrMandatoryName,
		},
		{
			name:    "should return error because CPF is mandatory",
			fields:  fields{Name: "Paula", CPF: "", Secret: "76t"},
			wantErr: ErrMandatoryCPF,
		},
		{
			name:    "should return error because CPF invalid",
			fields:  fields{Name: "Joao", CPF: "11895", Secret: "676"},
			wantErr: ErrInvalidCPF,
		},
		{
			name:    "should return error because format birth invalid",
			fields:  fields{Name: "Joao", CPF: "22244455566", Secret: "34n", BirthDate: "13-02-2000"},
			wantErr: ErrFormat,
		},
		{
			name:    "should return error because secret is mandatory",
			fields:  fields{Name: "Joao", CPF: "22244455566", Secret: ""},
			wantErr: ErrMandatorySecret,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &CreateRequest{
				Name:   tt.fields.Name,
				CPF:    tt.fields.CPF,
				Secret: tt.fields.Secret,
			}

			if err := a.Validate(); err != tt.wantErr {
				t.Errorf("CreateRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
