package option_test

import (
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
