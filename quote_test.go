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

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSymbol    = "TEST.PA"
	testTable     = "test_table"
	quoteErrorMsg = "quote_error_msg"
)

type DummyYApi struct {
	query string
	quote Quote
}

func (api *DummyYApi) Query(query string, responseObject interface{}) error {
	api.query = query
	quote := responseObject.(*QuotesResponse)
	quote.Query = QuotesQuery{
		Count: 1,
		Results: QuotesResults{
			Quote: api.quote,
		},
	}
	return nil
}

type ErrorYAPI struct {
	msg string
}

func (api *ErrorYAPI) Query(query string, responseObject interface{}) error {
	return errors.New(api.msg)
}

func TestGetQuote(t *testing.T) {
	expectedQueryString := fmt.Sprintf("select * from %s where symbol = \"%s\"", testTable, testSymbol)
	expectedQuote := Quote{
		Symbol:                     testSymbol,
		Name:                       "Test Name",
		LastTradePriceOnly:         11.1,
		FiftydayMovingAverage:      22.2,
		TwoHundreddayMovingAverage: 33.3,
		Volume: 44,
	}
	api := &DummyYApi{
		quote: expectedQuote,
	}
	quotes := &Quotes{
		IYApi: api,
		table: testTable,
	}
	quote, err := quotes.GetQuote(testSymbol)
	assert.Nil(t, err)
	assert.Equal(t, expectedQueryString, api.query)
	assert.Exactly(t, expectedQuote, *quote)
}

func TestGetQuoteError(t *testing.T) {
	quotes := &Quotes{
		IYApi: &ErrorYAPI{msg: quoteErrorMsg},
	}
	_, err := quotes.GetQuote(testSymbol)
	assert.NotNil(t, err)
}
