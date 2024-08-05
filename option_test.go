package option_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cospectrum/option"
	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	t.Parallel()

	var none option.Option[int]
	assert.True(t, none == nil)

	assert.True(t, option.None[int]() == nil)
	assert.True(t, option.Some(0) != nil)

	// :(
	opt := option.Option[int]{}
	assert.True(t, opt != nil)
}

func TestIsNone(t *testing.T) {
	t.Parallel()

	var none option.Option[int]
	assert.True(t, none.IsNone())

	assert.True(t, option.None[int]().IsNone())

	some := option.Some(0)
	assert.False(t, some.IsNone())

	opt := option.Option[int]{}
	assert.True(t, opt.IsNone())
}

func TestIsSome(t *testing.T) {
	t.Parallel()

	var none option.Option[int]
	assert.False(t, none.IsSome())

	assert.False(t, option.None[int]().IsSome())

	some := option.Some(0)
	assert.True(t, some.IsSome())

	opt := option.Option[int]{}
	assert.False(t, opt.IsSome())
}

func TestTake(t *testing.T) {
	t.Parallel()
	const val = 3

	opt := option.Some(val)
	other := opt.Take()

	assert.True(t, opt.IsNone())
	assert.Panics(t, func() {
		_ = opt.Unwrap()
	})

	assert.True(t, other.Unwrap() == val)
}

func TestClone(t *testing.T) {
	t.Parallel()
	const val = 3

	opt := option.Some(val)

	clone := opt.Clone()
	_ = clone.Take()
	assert.True(t, clone.IsNone())

	assert.Equal(t, opt.Unwrap(), val)
}

func TestUnwrapOrDefault(t *testing.T) {
	t.Parallel()
	const (
		val  = -1
		zero = 0
	)

	assert.Equal(t, zero, option.None[int]().UnwrapOrDefault())
	assert.Equal(t, val, option.Some(val).UnwrapOrDefault())
}

func TestReadme(t *testing.T) {
	divide := func(numerator, denominator float64) option.Option[float64] {
		if denominator == 0.0 {
			return nil // same as option.None[float64]()
		}
		return option.Some(numerator / denominator)
	}

	// The return value of the function is an option
	result := divide(2.0, 3.0)

	// Pattern match to retrieve the value
	err := option.Match(result,
		func(val float64) error {
			fmt.Printf("Result: %v\n", val)
			return nil
		},
		func() error {
			return errors.New("Cannot divide by 0")
		})
	if err != nil {
		panic(err)
	}
}
