# Env

Golang Get Environment Variables Package

![Build Status](https://github.com/nasermirzaei89/env/actions/workflows/build.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/env)](https://goreportcard.com/report/github.com/nasermirzaei89/env)
[![Codecov](https://codecov.io/gh/nasermirzaei89/env/branch/master/graph/badge.svg)](https://codecov.io/gh/nasermirzaei89/env)
[![Go Reference](https://pkg.go.dev/badge/github.com/nasermirzaei89/env.svg)](https://pkg.go.dev/github.com/nasermirzaei89/env)
[![awesome-go](https://awesome.re/badge.svg)](https://github.com/avelino/awesome-go#configuration)
[![License](https://img.shields.io/github/license/nasermirzaei89/env)](https://raw.githubusercontent.com/nasermirzaei89/env/master/LICENSE)

## Install

```sh
go get github.com/nasermirzaei89/env
```

## Sample Usage

### With default value

```go
package main

import (
	"fmt"
	"time"

	"github.com/nasermirzaei89/env"
)

func main() {
	b := env.GetBool("A", true)
	fmt.Println(b) // true (default)

	f := env.GetFloat64("B", 14.5)
	fmt.Println(f) // 14.5 (default)

	i := env.GetInt("C", 12)
	fmt.Println(i) // 12 (default)

	s := env.GetString("B", "hi")
	fmt.Println(s) // hi (default)

	d := env.GetDuration("D", 5*time.Second)
	fmt.Println(d) // 5s (default)

	// With generics

	b2 := env.Get("A", true)
	fmt.Println(b2) // true (default)

	f2 := env.Get("B", 14.5)
	fmt.Println(f2) // 14.5 (default)

	i2 := env.Get("C", 12)
	fmt.Println(i2) // 12 (default)

	s2 := env.Get("B", "hi")
	fmt.Println(s2) // hi (default)

	d2 := env.Get("D", 5*time.Second)
	fmt.Println(d2) // 5s (default)
}
```

### Force setting environment

```go
package main

import (
	"fmt"
	"time"

	"github.com/nasermirzaei89/env"
)

func main() {
	s := env.MustGetString("HOME")
	fmt.Println(s) // /Users/nasermirzaei89

	s = env.MustGetString("NEW") // panics

	d := env.MustGetDuration("TIMEOUT")
	fmt.Println(d) // e.g. 30s

	d = env.MustGetDuration("NEW") // panics

	// Generics

	s2 := env.MustGet[string]("HOME")
	fmt.Println(s2) // /Users/nasermirzaei89

	s2 = env.MustGet[string]("NEW") // panics

	d2 := env.MustGet[time.Duration]("TIMEOUT")
	fmt.Println(d2) // e.g. 30s

	d2 = env.MustGet[time.Duration]("NEW") // panics
}
```

### Lookup with structured errors

`LookupXXX` returns the value and an `error`. Use `errors.As` to distinguish between a missing key (`NotSetError`) and an unparseable value (`InvalidValueError`).

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/nasermirzaei89/env"
)

func main() {
	// Typed lookup
	v, err := env.LookupInt("PORT")
	if err != nil {
		var notSet env.NotSetError
		var invalid env.InvalidValueError

		switch {
		case errors.As(err, &notSet):
			fmt.Printf("%q is not set\n", notSet.Key)
		case errors.As(err, &invalid):
			fmt.Printf("%q has invalid value %q: %v\n", invalid.Key, invalid.Value, invalid.Err)
		}
	} else {
		fmt.Println("PORT =", v)
	}

	// Generic lookup
	s, err := env.Lookup[string]("HOME")
	if err == nil {
		fmt.Println("HOME =", s)
	}
}
```

## Contributing

You can submit a [new issue](https://github.com/nasermirzaei89/env/issues/new) in
GitHub [issues](https://github.com/nasermirzaei89/env/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create
a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get
merged into the main package.
