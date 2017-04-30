// Copyright 2017 Clément Bizeau
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

import "fmt"

const (
	quotesTable = "yahoo.finance.quotes"
)

// Quote represents a quote on yahoo api
type Quote struct {
	Symbol                     string  `json:"symbol"`
	Name                       string  `json:"Name"`
	LastTradePriceOnly         float32 `json:"LastTradePriceOnly,string"`
	FiftydayMovingAverage      float32 `json:"FiftydayMovingAverage,string"`
	TwoHundreddayMovingAverage float32 `json:"TwoHundreddayMovingAverage,string"`
	Volume                     int32   `json:"Volume,string"`
}

// QuotesResults represents the results of an history query
type QuotesResults struct {
	Quote Quote `json:"quote"`
}

// QuotesQuery contains the data of a quotes query
type QuotesQuery struct {
	Count   int           `json:"count"`
	Results QuotesResults `json:"results"`
}

// QuotesResponse is the response of the quotes query
type QuotesResponse struct {
	Query QuotesQuery `json:"query"`
}

// QuotesAPI allows user to Access yahoo finance quotes api
type QuotesAPI interface {
	GetQuote(symbol string) (*Quote, error)
}

// Quotes allows user to access yahoo finance quotes api
type Quotes struct {
	IYApi
	table string
}

// NewQuotes creates a new Quotes api object
//
// Return the new quotes api object
func NewQuotes() QuotesAPI {
	return &Quotes{
		table: quotesTable,
		IYApi: NewYApi(),
	}
}

// GetQuote retrieves the quote values from yahoo yql
//
//  GetQuote("CW8.PA")
//
// Returns the quote values
func (quotes *Quotes) GetQuote(symbol string) (*Quote, error) {
	query := fmt.Sprintf("select * from %s where symbol = \"%s\"", quotes.table, symbol)
	var res QuotesResponse
	if err := quotes.Query(query, &res); err != nil {
		return nil, err
	}
	return &res.Query.Results.Quote, nil
}
