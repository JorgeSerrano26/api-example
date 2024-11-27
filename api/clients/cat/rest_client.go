package cat

import (
	"encoding/json"
	"github.com/JorgeSerrano26/api-example/api/entities"
	"net/http"
	"net/url"
	"time"
)

const (
	factPath = "/fact"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "catfact.ninja",
}

var catClient = initClient()

type catRestClient struct{}

func (c catRestClient) GetFact() (entities.CatFact, error) {
	req, err := http.NewRequest(http.MethodGet, c.buildURL(factPath), nil)

	if err != nil {
		return entities.CatFact{}, err
	}

	res, err := catClient.Do(req)

	if err != nil {
		return entities.CatFact{}, err
	}

	defer res.Body.Close()

	var catFact entities.CatFact

	if err := json.NewDecoder(res.Body).Decode(&catFact); err != nil {
		return entities.CatFact{}, err
	}

	return catFact, nil
}

func (c catRestClient) buildURL(path string) string {
	endpoint := baseURL.ResolveReference(&url.URL{Path: path})
	return endpoint.String()
}

func initClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
