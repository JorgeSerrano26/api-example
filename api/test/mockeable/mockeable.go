package mockeable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mockeable interface {
	GetFuncControls() []*CallsFuncControl
	CleanUp()
	Use()
}

func CleanUpAndAssertControls(t *testing.T, mock Mockeable) {
	defer mock.CleanUp()
	for _, control := range mock.GetFuncControls() {
		if !control.IgnoreCallsAssertion {
			assert.Equal(t, control.ExpectedCalls, control.funcCalls,
				fmt.Sprintf("expected calls for func %s does not match actual", control.funcName),
			)
		}
	}
}
