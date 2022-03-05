package promise_test

import (
	"fmt"
	"testing"

	"generic-packages/pkg/promise"
	"generic-packages/pkg/result"

	"github.com/stretchr/testify/assert"
)

func TestLazy(t *testing.T) {
	a := assert.New(t)

	calls := 0
	p := promise.Lazy(func() result.Result[int] { calls++; return result.OK(1) })

	a.Zero(calls)
	a.Equal(result.OK(1), p.Await())
	a.Equal(1, calls)
	a.Equal(result.OK(1), p.Await())
	a.Equal(1, calls)
}

func TestEager(t *testing.T) {
	a := assert.New(t)

	calls := 0
	p := promise.Eager(func() result.Result[int] { calls++; return result.OK(1) })

	a.Equal(1, calls)
	a.Equal(result.OK(1), p.Await())
	a.Equal(1, calls)
	a.Equal(result.OK(1), p.Await())
	a.Equal(1, calls)
}

func TestMap(t *testing.T) {
	a := assert.New(t)
	a.Equal(result.OK(1), promise.Map(Identity[int], promise.Resolved(result.OK(1))).Await())
	a.Equal(result.Error[string](fmt.Errorf("some error")), promise.Map(assertNotCalled[int, string](a), promise.Resolved(result.Error[int](fmt.Errorf("some error")))).Await())
}

func Identity[T any](t T) T {
	return t
}

func assertNotCalled[T, U any](a *assert.Assertions) func(T) U {
	return func(t T) (u U) {
		a.Fail("func not expected to be called", "func called with %v", t)
		return
	}
}
