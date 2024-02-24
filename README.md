# Env

Golang Get Environment Variables Package

![Build Status](https://github.com/nasermirzaei89/env/actions/workflows/build.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/env)](https://goreportcard.com/report/github.com/nasermirzaei89/env)
[![Codecov](https://codecov.io/gh/nasermirzaei89/env/branch/master/graph/badge.svg)](https://codecov.io/gh/nasermirzaei89/env)
[![Go Reference](https://pkg.go.dev/badge/github.com/nasermirzaei89/env.svg)](https://pkg.go.dev/github.com/nasermirzaei89/env)
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

	// Generics

	b2 := env.Get("A", true)
	fmt.Println(b2) // true (default)

	f2 := env.Get("B", 14.5)
	fmt.Println(f2) // 14.5 (default)

	i2 := env.Get("C", 12)
	fmt.Println(i2) // 12 (default)

	s2 := env.Get("B", "hi")
	fmt.Println(s2) // hi (default)
}
```

### Force setting environment

```go
package main

import (
	"fmt"

	"github.com/nasermirzaei89/env"
)

func main() {
	s := env.MustGetString("HOME")
	fmt.Println(s) // /Users/nasermirzaei89

	s = env.MustGetString("NEW") // panics

	// Generics

	s2 := env.MustGet[string]("HOME")
	fmt.Println(s2) // /Users/nasermirzaei89

	s2 = env.MustGet[string]("NEW") // panics
}
```

## Contributing

You can submit a [new issue](https://github.com/nasermirzaei89/env/issues/new) in
GitHub [issues](https://github.com/nasermirzaei89/env/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create
a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get
merged into the main package.
