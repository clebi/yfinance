// Copyright 2016 Clément Bizeau
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package finance

import (
	"fmt"
	"time"
)

const (
	historyTable = "yahoo.finance.historicaldata"
)

// Stock is an object containing stock values
type Stock struct {
	Date   YTime   `json:"Date,string"`
	Open   float32 `json:"Open,string"`
	High   float32 `json:"High,string"`
	Low    float32 `json:"Low,string"`
	Close  float32 `json:"Close,string"`
	Volume int     `json:"Volume,string"`
	Symbol string  `json:"Symbol"`
}

// HistoryResults represents the results of an history query
type HistoryResults struct {
	Stocks []Stock `json:"quote"`
}

// HistoryQuery contains the data of an history query
type HistoryQuery struct {
	Count   int            `json:"count"`
	Results HistoryResults `json:"results"`
}

// HistoryResponse is the response of the history query
type HistoryResponse struct {
	Query HistoryQuery `json:"query"`
}

// HistoryAPI allows to retrieve stocks history
type HistoryAPI interface {
	GetHistory(symbol string, start time.Time, end time.Time) ([]Stock, error)
}

// History is the object to manage history api
type History struct {
	IYApi
	table string
}

// NewHistory creates a new History api object
//
// Return the new history api object
func NewHistory() HistoryAPI {
	return &History{
		table: historyTable,
		IYApi: NewYApi(),
	}
}

// NewHistoryTest create an new history api object for test
//
// Returns the new history api test object
func NewHistoryTest(mockHistory IYApi) HistoryAPI {
	return &History{
		table: historyTable,
		IYApi: mockHistory,
	}
}

// GetHistory can be used to get the history values of a finance stock
//
//  history.GetHistory("GOOG", startDate, endDate)
//
// Returns an array with the stock values for the corresponding period
func (history *History) GetHistory(symbol string, start time.Time, end time.Time) ([]Stock, error) {
	query := fmt.Sprintf("select * from %s where symbol = \"%s\" and "+
		"startDate = \"%s\" and endDate = \"%s\"", historyTable, symbol, start.Format(DateFormat), end.Format(DateFormat))
	var res HistoryResponse
	err := history.Query(query, &res)
	if err != nil {
		return nil, err
	}
	return res.Query.Results.Stocks, nil
}
