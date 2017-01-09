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
)

func TestMarshalJSON(t *testing.T) {
	now := YTime{time.Now()}
	timeBytes, err := now.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\""+now.Format(DateFormat)+"\"", string(timeBytes))
}

func TestUnmarshalJSON(t *testing.T) {
	testDate, err := time.Parse(DateFormat, "2016-12-03")
	if err != nil {
		t.Fatal(err)
	}
	now := YTime{testDate}
	timeBytes := []byte(now.Format(DateFormat))
	var date YTime
	date.UnmarshalJSON(timeBytes)
	assert.Equal(t, now, date)
}

func TestUnmarshalJSONError(t *testing.T) {
	timeBytes := []byte("error")
	var date YTime
	err := date.UnmarshalJSON(timeBytes)
	t.Log(err)
	assert.NotNil(t, err)
}
