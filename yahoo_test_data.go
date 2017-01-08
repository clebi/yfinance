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

const (
	testQuerySelect   = "select test"
	testQueryLength   = 1
	testQueryResponse = `{
 "query": {
  "count": 1,
  "created": "2017-01-08T13:02:56Z",
  "lang": "en-US",
  "results": {
   "quote": [{
    "Symbol": "TEST",
    "Date": "2016-12-01",
    "Open": "1.1",
    "High": "1.2",
    "Low": "1.3",
    "Close": "1.4",
    "Volume": "15",
    "Adj_Close": "1.6"
   }]
  }
 }
}`
	testQueryAPIErrorResponse = `{
  "error": {
    "lang": "en-US",
    "description": "error"
  }
}`
)
