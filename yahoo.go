package finance

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	apiURL = "https://query.yahooapis.com/v1/public/yql"
	// DateFormat for yahoo finance api
	DateFormat = "2006-01-02"
)

// YApiErrorContent represents the content a yahoo api error
type YApiErrorContent struct {
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

// YApiError represents a yahoo api error
type YApiError struct {
	Content YApiErrorContent `json:"error"`
}

func (err YApiError) Error() string {
	return err.Content.Description
}

// IYApi is the interface for the yahoo api
type IYApi interface {
	Query(query string, responseObject interface{}) error
}

// YApi is the yahoo finance api
type YApi struct {
	apiURL string
	http   *http.Client
}

// NewYApi create a new yahoo finance api error
//
// Returns the new api object
func NewYApi() *YApi {
	return &YApi{
		apiURL: apiURL,
		http:   &http.Client{},
	}
}

// Query execute a query on yahoo finance api
// responseObject will be used to store the api response
//
//  api.Query("select * from table where startDate = \"2016-01-01\"", responseObject)
//
// Returns an error if something went wrong
func (api *YApi) Query(query string, responseObject interface{}) error {
	u, err := url.Parse(api.apiURL)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("q", query)
	q.Set("format", "json")
	q.Set("env", "store://datatables.org/alltableswithkeys")
	u.RawQuery = q.Encode()
	resp, err := api.http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var yerr YApiError
		err = json.NewDecoder(resp.Body).Decode(&yerr)
		if err != nil {
			return err
		}
		return yerr
	}
	err = json.NewDecoder(resp.Body).Decode(responseObject)
	if err != nil {
		return err
	}
	return nil
}
