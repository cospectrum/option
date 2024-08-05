package option_test

import (
	"encoding/json"
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

	type U struct {
		Num option.Option[int] `json:"num"`
	}

	var u U
	err = json.Unmarshal([]byte(`{"num": null}`), &u)
	// => U{Num: option.None()}
	assert.NoError(t, err)
	assert.True(t, u.Num.IsNone())

	err = json.Unmarshal([]byte(`{}`), &u)
	// => U{Num: option.None()}
	assert.NoError(t, err)
	assert.True(t, u.Num.IsNone())

	err = json.Unmarshal([]byte(`{"num": 0}`), &u)
	// => U{Num: option.Some(0)}
	assert.NoError(t, err)
	assert.Equal(t, 0, u.Num.Unwrap())

	err = json.Unmarshal([]byte(`{"num": 3}`), &u)
	// => U{Num: option.Some(3)}
	assert.NoError(t, err)
	assert.Equal(t, 3, u.Num.Unwrap())
}

func TestJSON(t *testing.T) {
	type T struct {
		Num option.Option[int] `json:"num"`
	}
	var ty T

	s := `{"num": null}`
	err := json.Unmarshal([]byte(s), &ty)
	assert.NoError(t, err)
	assert.True(t, ty.Num.IsNone())

	s = `{}`
	err = json.Unmarshal([]byte(s), &ty)
	assert.NoError(t, err)
	assert.True(t, ty.Num.IsNone())

	s = `{"num": 0}`
	err = json.Unmarshal([]byte(s), &ty)
	assert.NoError(t, err)
	assert.Equal(t, 0, ty.Num.Unwrap())

	s = `{"num": 3}`
	err = json.Unmarshal([]byte(s), &ty)
	assert.NoError(t, err)
	assert.Equal(t, 3, ty.Num.Unwrap())

	id := func(ty T) T {
		b, err := json.Marshal(ty)
		assert.NoError(t, err)

		var out T
		assert.NoError(t, json.Unmarshal(b, &out))
		return out
	}

	vals := []T{
		{},
		{Num: nil},
		{Num: option.None[int]()},
		{Num: option.Some(0)},
		{Num: option.Some(3)},
	}
	for _, val := range vals {
		newVal := id(val)
		assert.Equal(t, val, newVal)

		if val.Num.IsSome() {
			assert.Equal(t, val.Num.Unwrap(), newVal.Num.Unwrap())
			continue
		}

		assert.True(t, val.Num.IsNone())
		assert.True(t, newVal.Num.IsNone())
	}
}
