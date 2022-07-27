package datastore

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Store interface {
	Connect(ctx context.Context) error
}

type GsheetProvider struct {
	Token   string
	service *sheets.Service
}

func NewGsheetProvider(token string) *GsheetProvider {
	return &GsheetProvider{
		Token: token,
	}
}

func (s *GsheetProvider) Connect(ctx context.Context) error {
	sheetService, err := sheets.NewService(ctx, option.WithAPIKey(s.Token))
	if err != nil {
		return err
	}
	s.service = sheetService
	return nil
}

func (gp *GsheetProvider) Get(sheetId string) {
	readRange := "PRECIOS TOTALES!A3:O"
	resp, err := gp.service.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Item, Nombre, Precio:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s, %s\n", row[0], row[1], row[2])
		}
	}
}

// Range: A2:O73
