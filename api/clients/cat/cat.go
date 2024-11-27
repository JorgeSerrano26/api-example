package cat

import "github.com/JorgeSerrano26/api-example/api/entities"

type Client interface {
	GetFact() (entities.CatFact, error)
}

var (
	defaultInstance Client = catRestClient{}
	Instance        Client = defaultInstance
)
