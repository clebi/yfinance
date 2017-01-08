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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	selectStr = "select * from yahoo.finance.historicaldata where symbol = \"TEST\" and " +
		"startDate = \"2016-12-01\" and endDate = \"2016-12-31\""
)

type MockYApi struct {
	mock.Mock
}

func (api *MockYApi) Query(query string, responseObject interface{}) error {
	args := api.Called(query, responseObject)
	return args.Error(0)
}

func TestGetHistory(t *testing.T) {
	stocks := []Stock{{Open: 1.1, High: 2.2, Low: 3.3, Close: 4.4, Volume: 999, Symbol: "TEST"}}
	var argObject HistoryResponse
	mockYApi := &MockYApi{}
	mockYApi.On("Query", selectStr, &argObject).
		Return(nil).
		Run(func(args mock.Arguments) {
			response := args.Get(1).(*HistoryResponse)
			response.Query = HistoryQuery{
				Count: 1,
				Results: HistoryResults{
					Stocks: stocks,
				},
			}
		})
	history := NewHistoryTest(mockYApi)
	start, _ := time.Parse(DateFormat, "2016-12-01")
	end, _ := time.Parse(DateFormat, "2016-12-31")
	res, err := history.GetHistory("TEST", start, end)
	assert.Nil(t, err)
	assert.Exactly(t, stocks, res)
	mockYApi.AssertExpectations(t)
}

func TestGetHistoryError(t *testing.T) {
	errorStr := "error"
	var argObject HistoryResponse
	mockYApi := &MockYApi{}
	mockYApi.On("Query", selectStr, &argObject).Return(YApiError{
		Content: YApiErrorContent{
			Lang:        "en-US",
			Description: "error",
		},
	})
	history := NewHistoryTest(mockYApi)
	start, _ := time.Parse(DateFormat, "2016-12-01")
	end, _ := time.Parse(DateFormat, "2016-12-31")
	_, err := history.GetHistory("TEST", start, end)
	assert.NotNil(t, err)
	assert.Equal(t, errorStr, err.Error())
	mockYApi.AssertExpectations(t)
}
