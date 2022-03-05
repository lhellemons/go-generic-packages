package iterator_test

import (
	"testing"

	"generic-packages/pkg/iterator"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	AssertEqualCollected(t, iterator.Static(2, 4, 6), iterator.Map(TimesTwo, iterator.Static(1, 2, 3)))
	AssertEqualCollected(t, iterator.Empty[int](), iterator.Map(TimesTwo, iterator.Empty[int]()))
}

func TimesTwo(i int) int {
	return i * 2
}

func AssertEqualCollected[T any](t *testing.T, expected iterator.Iterator[T], actual iterator.Iterator[T]) {
	assert.Equal(t, iterator.Collect(expected), iterator.Collect(actual))
}

// TODO add test for rest of funcs
