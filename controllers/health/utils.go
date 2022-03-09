package health

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/burntcarrot/apollo/entity/health"
)

func getResults(services []health.Domain) ([]Result, error) {
	var results []Result

	for _, srv := range services {
		// check health of service
		data, err := http.Get(srv.URI)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(data.Body)
		if err != nil {
			return nil, errors.New("failed to read from service")
		}
		defer data.Body.Close()

		resp := Response{}
		json.Unmarshal([]byte(body), &resp)
		results = append(results, resp.Results...)
	}

	return results, nil
}
