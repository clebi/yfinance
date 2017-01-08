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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Equal(t, testQuerySelect, q.Get(apiQueryKey))
		assert.Equal(t, apiFormat, q.Get(apiFormatKey))
		assert.Equal(t, apiEnv, q.Get(apiEnvKey))
		fmt.Fprint(w, testQueryResponse)
	}))
	defer ts.Close()
	api := NewYApiTest(ts.URL)
	var responseObject HistoryResponse
	err := api.Query(testQuerySelect, &responseObject)
	assert.Nil(t, err)
	assert.Equal(t, testQueryLength, responseObject.Query.Count)
	assert.Len(t, responseObject.Query.Results.Stocks, testQueryLength)
	stock := responseObject.Query.Results.Stocks[0]
	assert.EqualValues(t, 1.1, stock.Open)
	assert.EqualValues(t, 1.2, stock.High)
	assert.EqualValues(t, 1.3, stock.Low)
	assert.EqualValues(t, 1.4, stock.Close)
	assert.EqualValues(t, 15, stock.Volume)
}

func TestQueryAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, testQueryAPIErrorResponse)
	}))
	defer ts.Close()
	api := NewYApiTest(ts.URL)
	var responseObject HistoryResponse
	err := api.Query("error", &responseObject)
	assert.NotNil(t, err)
	var yapiError YApiError
	assert.IsType(t, yapiError, err)
}
