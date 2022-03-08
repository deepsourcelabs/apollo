package health

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/burntcarrot/apollo/entity/health"
)

func getResults(services []health.Domain) ([]health.Result, error) {
	var results []health.Result

	for _, srv := range services {
		// check health of service
		resp, err := http.Get(srv.URI)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("failed to read from service")
		}
		defer resp.Body.Close()

		respNew := health.Response{}
		json.Unmarshal([]byte(body), &respNew)
		results = append(results, respNew.Results...)
	}

	return results, nil
}
