package health

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/burntcarrot/apollo/entity/health"
)

// getResults scrapes health status from services.
func getResults(services []health.Domain) ([]Result, error) {
	var results []Result

	for _, srv := range services {
		// fetch health status
		data, err := http.Get(srv.URI)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(data.Body)
		if err != nil {
			return nil, errors.New("failed to read from service")
		}
		defer data.Body.Close()

		// unmarshal response body into struct
		resp := Response{}
		err = json.Unmarshal([]byte(body), &resp)
		if err != nil {
			return nil, errors.New("failed to unmarshal response body")
		}
		results = append(results, resp.Results...)
	}

	return results, nil
}
