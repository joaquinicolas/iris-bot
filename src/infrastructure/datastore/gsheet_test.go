package datastore

import (
	"context"
	"testing"
)

func Test_Gsheet_Connect(t *testing.T) {
	token := "test_token"
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Test_Gsheet_Connect succesfully", wantErr: false},
		{name: "Test_Gsheet_Connect Get error", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gsheet := NewGsheetProvider(token)
			if gsheet == nil {
				t.Errorf("Expected gsheet, got nil")
			}

			ctx := context.Background()
			if err := gsheet.Connect(ctx); err != nil {
				if tt.wantErr {
					t.Errorf("Gsheet.Connect() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_Gsheet_Get(t *testing.T) {
	tests := []struct {
		name      string
		readRange string
		sheetId   string
		wantErr   bool
	}{
		{name: "Test_Gsheet_Get succesfully", wantErr: false, readRange: "PRECIOS TOTALES!A3:O80", sheetId: "1x3UNLdLbmnl0d65fEmH2us4Xj15Xj9ZQ-UNM5SXFVDg"},
		//{name: "Test_Gsheet_Get Get error", wantErr: true, readRange: "test", sheetId: "1"},
	}

	gsheet := NewGsheetProvider("AIzaSyB05LD8IrWbvmQLR6a0dUicsoePnGirsH8")
	if gsheet == nil {
		t.Errorf("Expected gsheet, got nil")
	}

	ctx := context.Background()
	gsheet.Connect(ctx)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gsheet.Get(tt.sheetId, tt.readRange)
			if err != nil && !tt.wantErr {
				t.Errorf("Gsheet.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("Expected resp, got nil")
				return
			}

			if len(resp) == 0 && !tt.wantErr {
				t.Errorf("Expected resp, got nil")
				return
			}

			for _, v := range resp {
				if v.Id == "" {
					t.Errorf("Expected Id, got nil")
				}
				if v.Name == "" {
					t.Errorf("Expected Name, got nil")
				}
			}
		})
	}
}
