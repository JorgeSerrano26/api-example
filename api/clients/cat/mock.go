package cat

import (
	"github.com/JorgeSerrano26/api-example/api/entities"
	"github.com/JorgeSerrano26/api-example/api/test/mockeable"
)

type MockClient struct {
	GetFactFunc func() (entities.CatFact, error)

	GetFactControl mockeable.CallsFuncControl
}
