package config

import (
	"testing"
)

func TestRules_Validate(t *testing.T) {
	type fields struct {
		WinPoints  int
		DrawPoints int
		LossPoints int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "when win points is less than 0",
			fields: fields{
				WinPoints:  -1,
				DrawPoints: 0,
				LossPoints: 0,
			},
			wantErr: true,
		}, {
			name: "when draw points is less than 0",
			fields: fields{
				WinPoints:  0,
				DrawPoints: -1,
				LossPoints: 0,
			},
			wantErr: true,
		}, {
			name: "when loss points is less than 0",
			fields: fields{
				WinPoints:  0,
				DrawPoints: 0,
				LossPoints: -1,
			},
			wantErr: true,
		}, {
			name: "when all points are greater than or equal to 0",
			fields: fields{
				WinPoints:  0,
				DrawPoints: 0,
				LossPoints: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rules{
				WinPoints:  tt.fields.WinPoints,
				DrawPoints: tt.fields.DrawPoints,
				LossPoints: tt.fields.LossPoints,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
