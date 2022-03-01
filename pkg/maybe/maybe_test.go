package maybe_test

import (
	"strconv"
	"testing"

	"generic-utils/pkg/maybe"

	"github.com/stretchr/testify/assert"
)

func TestMap_JustOne(t *testing.T) {
	assert.Equal(t,
		maybe.Just("1"),
		maybe.Map(strconv.Itoa, maybe.Just(1)))
}

func TestMap_Nothing(t *testing.T) {
	assert.Equal(t,
		maybe.Nothing[string](),
		maybe.Map(strconv.Itoa, maybe.Nothing[int]()))
}

func TestFlatMap_JustOne(t *testing.T) {
	assert.Equal(t,
		maybe.Just("1"),
		maybe.FlatMap(ItoMaybeA, maybe.Just(1)))
}
func TestFlatMap_Nothing(t *testing.T) {
	assert.Equal(t,
		maybe.Nothing[string](),
		maybe.FlatMap(ItoNothing, maybe.Just(1)))
}

func ItoMaybeA(i int) maybe.Maybe[string] {
	return maybe.Just(strconv.Itoa(i))
}

func ItoNothing(i int) maybe.Maybe[string] {
	return maybe.Nothing[string]()
}
