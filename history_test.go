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
	stocks := []Stock{Stock{Open: 1.1, High: 2.2, Low: 3.3, Close: 4.4, Volume: 999, Symbol: "TEST"}}
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
