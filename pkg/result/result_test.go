package result_test

import (
	"fmt"
	"testing"

	"generic-utils/pkg/result"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t, result.OK(2), result.Map(TimesTwo, result.OK(1)))
	assert.Equal(t, result.OK("a"), result.Map(Identity[string], result.OK("a")))

	assert.Equal(t, result.Error[int](fmt.Errorf("some error")), result.Map(TimesTwo, result.Error[int](fmt.Errorf("some error"))))
	assert.Equal(t, result.Error[string](fmt.Errorf("some error")), result.Map(Identity[string], result.Error[string](fmt.Errorf("some error"))))
}

func Identity[T any](t T) T {
	return t
}

func TimesTwo(i int) int {
	return i * 2
}

func TestFlatmap(t *testing.T) {
	assert.Equal(t, result.OK(5.0), result.Flatmap(Divide10By, result.OK(2)))
	assert.Equal(t, result.Error[float64](fmt.Errorf("invalid input")), result.Flatmap(Divide10By, result.Error[int](fmt.Errorf("invalid input"))))
	assert.Equal(t, result.Error[float64](fmt.Errorf("cannot divide by zero")), result.Flatmap(Divide10By, result.OK(0)))
}

func Divide10By(n int) result.Result[float64] {
	if n == 0 {
		return result.Error[float64](fmt.Errorf("cannot divide by zero"))
	}
	return result.OK(float64(10.0 / n))
}
