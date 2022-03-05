package list_test

import (
	"testing"

	"generic-packages/pkg/list"

	"github.com/stretchr/testify/assert"
)

func TestList_Empty(t *testing.T) {
	assert.Equal(t, []int{}, list.Empty[int]().Items())
	assert.Equal(t, []string{}, list.Empty[string]().Items())
}

func TestList_FromSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, list.Slice([]int{1, 2, 3}).Items())
	assert.Equal(t, []string{"foo", "bar"}, list.Slice([]string{"foo", "bar"}).Items())
}

func TestMap(t *testing.T) {
	assert.Equal(t,
		list.Static(2, 4, 6),
		list.Map(func(i int) int { return i * 2 }, list.Static(1, 2, 3)))
}
