package option

import (
	"bytes"
	"encoding/json"
)

type Option[T any] []T

func Some[T any](val T) Option[T] {
	return Option[T]{val}
}

func None[T any]() Option[T] {
	return nil
}

// Returns true if the option has a value.
func (opt Option[T]) IsSome() bool {
	return !opt.IsNone()
}

func Match[T any, U any](opt Option[T], some func(T) U, none func() U) U {
	if opt.IsSome() {
		return some(opt.Unwrap())
	}
	return none()
}

// Returns true if the option has a value and the value matches a predicate.
func (opt Option[T]) IsSomeAnd(f func(T) bool) bool {
	if opt.IsSome() {
		return f(opt.Unwrap())
	}
	return false
}

// Returns true if the option has no value.
func (opt Option[T]) IsNone() bool {
	return len(opt) == 0
}

// Returns the contained value.
// Panics if there is no value with a custom panic message.
func (opt Option[T]) Expect(msg string) T {
	if opt.IsNone() {
		panic(msg)
	}
	return opt[0]
}

// Returns the contained value.
// Panics if there is no value.
func (opt Option[T]) Unwrap() T {
	const msg = "called `Option.Unwrap()` on a `None` value"
	return opt.Expect(msg)
}

func (opt Option[T]) UnwrapOr(defaultVal T) T {
	return Match(opt,
		func(val T) T {
			return val
		}, func() T {
			return defaultVal
		})
}

// Returns the contained value or computes it from a function.
func (opt Option[T]) UnwrapOrElse(f func() T) T {
	return Match(opt, func(val T) T {
		return val
	}, func() T {
		return f()
	})
}

// Returns the contained value or a default.
func (opt Option[T]) UnwrapOrDefault() T {
	return Match(opt, func(val T) T {
		return val
	}, func() T {
		var zero T
		return zero
	})
}

// Takes the value out of the option, leaving a None in its place.
func (opt *Option[T]) Take() Option[T] {
	if opt.IsSome() {
		val := opt.Unwrap()
		*opt = None[T]()
		return Some(val)
	}
	return nil
}

func (opt Option[T]) Clone() Option[T] {
	if opt.IsSome() {
		return Some(opt.Unwrap())
	}
	return None[T]()
}

var (
	_ json.Marshaler   = Option[any]{}
	_ json.Unmarshaler = &Option[any]{}
)

var jsonNull = []byte("null")

func (opt Option[T]) MarshalJSON() ([]byte, error) {
	if opt.IsNone() {
		return jsonNull, nil
	}
	marshal, err := json.Marshal(opt.Unwrap())
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (opt *Option[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, jsonNull) {
		*opt = None[T]()
		return nil
	}
	var val T
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	*opt = Some(val)
	return nil
}
