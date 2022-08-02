package mocks

import (
	"github.com/stretchr/testify/mock"
)

type SpreadsheetsValuesGetCall struct {
	mock.Mock
}

type SpreadsheetsValuesService struct {
	mock.Mock
}

type SpreadsheetsService struct {
	mock.Mock
	Values SpreadsheetsValuesService
}

type Service struct {
	mock.Mock
	Spreadsheets *SpreadsheetsService
}

func (ss *SpreadsheetsService) Get(sheetId string, readRange string) *SpreadsheetsValuesGetCall {
	if sheetId == "" {
		return nil
	}
	if readRange == "" {
		return nil
	}
	args := ss.Called(sheetId, readRange)
	return args.Get(0).(*SpreadsheetsValuesGetCall)
}
