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
	"bytes"
	"strings"
	"time"
)

// YTime represents a date in yahoo finance api
type YTime struct {
	time.Time
}

// MarshalJSON encode a date from yahoo api
func (t *YTime) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	buffer.WriteString(t.Format(DateFormat))
	buffer.WriteString("\"")
	return []byte(buffer.Bytes()), nil
}

// UnmarshalJSON decode a date from yahoo api
func (t *YTime) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(DateFormat, strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}
