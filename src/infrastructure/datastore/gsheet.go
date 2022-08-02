package datastore

import (
	"context"
	"fmt"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Store interface {
	Connect(ctx context.Context) error
	Get(sheetId string, readRange string) any
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

func (gp *GsheetProvider) Get(sheetId string, readRange string) ([]entities.Product, error) {
	//readRange := "PRECIOS TOTALES!A3:O80"
	var products = make([]entities.Product, 0)
	resp, err := gp.service.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	} else {
		for _, row := range resp.Values {
			if row[0] != "" {
				products = append(products, entities.Product{
					Id:    fmt.Sprintf("%v", row[0]),
					Price: fmt.Sprintf("%v", row[2]),
					Name:  fmt.Sprintf("%v", row[1]),
				})
			}

			if row[4] != "" {
				products = append(products, entities.Product{
					Id:    fmt.Sprintf("%v", row[4]),
					Price: fmt.Sprintf("%v", row[6]),
					Name:  fmt.Sprintf("%v", row[5]),
				})
			}

			if row[8] != "" {
				products = append(products, entities.Product{
					Id:    fmt.Sprintf("%v", row[8]),
					Price: fmt.Sprintf("%v", row[10]),
					Name:  fmt.Sprintf("%v", row[9]),
				})
			}

			if row[12] != "" {
				products = append(products, entities.Product{
					Id:    fmt.Sprintf("%v", row[12]),
					Price: fmt.Sprintf("%v", row[14]),
					Name:  fmt.Sprintf("%v", row[13]),
				})
			}
		}
	}

	return products, nil
}
