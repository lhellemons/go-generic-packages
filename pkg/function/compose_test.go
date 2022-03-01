package function_test

import (
	"testing"

	"generic-utils/pkg/function"

	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	assert.Equal(t, true, function.Compose(Not, Not)(true))
	assert.Equal(t, 8, function.Compose(Add(3), Multiply(2))(1))
	assert.Equal(t, 5, function.Compose(Multiply(2), Add(3))(1))
}

func Not(b bool) bool {
	return !b
}

func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func Multiply(k int) func(int) int {
	return func(a int) int {
		return a * k
	}
}
