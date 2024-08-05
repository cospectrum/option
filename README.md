# option

[![github]](https://github.com/cospectrum/option)
[![goref]](https://pkg.go.dev/github.com/cospectrum/option)

[github]: https://img.shields.io/badge/github-cospectrum/option-8da0cb?logo=github
[goref]: https://pkg.go.dev/badge/github.com/cospectrum/option

Option type for golang

## Install
```sh
go get -u github.com/cospectrum/option
```
Requires Go version `1.22.0` or greater.

## Usage
```go
import (
	"fmt"

	"github.com/cospectrum/option"
)

func main()
	divide := func(numerator, denominator float64) option.Option[float64] {
		if denominator == 0.0 {
			return option.None[float64]()
		}
		return option.Some(numerator / denominator)
	}

	// The return value of the function is an option
	result := divide(2.0, 3.0)

	// Pattern match to retrieve the value
	result.Match(
		func(val float64) {
			fmt.Printf("Result: %v\n", val)
		},
		func() {
			fmt.Println("Cannot divide by 0")
		},
	)
}
```

### JSON
```go
type U struct {
	Num option.Option[int] `json:"num"`
}

var u U
json.Unmarshal([]byte(`{"num": null}`), &u) // => U{Num: option.None()}
json.Unmarshal([]byte(`{}`), &u) // => U{Num: option.None()}
json.Unmarshal([]byte(`{"num": 0}`), &u) // => U{Num: option.Some(0)}
json.Unmarshal([]byte(`{"num": 3}`), &u) // => U{Num: option.Some(3)}
```
