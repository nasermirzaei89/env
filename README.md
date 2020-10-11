# Env

Golang Get Environment Variables Package

[![Build Status](https://travis-ci.org/nasermirzaei89/env.svg?branch=master)](https://travis-ci.org/nasermirzaei89/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/env)](https://goreportcard.com/report/github.com/nasermirzaei89/env)
[![Codecov](https://codecov.io/gh/nasermirzaei89/env/branch/master/graph/badge.svg)](https://codecov.io/gh/nasermirzaei89/env)
[![GoDoc](https://godoc.org/github.com/nasermirzaei89/env?status.svg)](https://godoc.org/github.com/nasermirzaei89/env)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nasermirzaei89/env)](https://pkg.go.dev/github.com/nasermirzaei89/env)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/nasermirzaei89/env/master/LICENSE)

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

func main()  {
    var b  = env.GetBool("A", true)
    fmt.Println(b) // true (default)

    var f  = env.GetFloat("B", 14.5)
    fmt.Println(f) // 14.5 (default)

    var i  = env.GetInt("C", 12)
    fmt.Println(i) // 12 (default)

    var s  = env.GetString("B", "hi")
    fmt.Println(s) // hi (default)
}
```

### Force setting environment

```go
package main

import (
    "fmt"
    "github.com/nasermirzaei89/env"
)

func main()  {
    var s  = env.MustGetString("HOME")
    fmt.Println(s) // /Users/nasermirzaei89

    s  = env.MustGetString("NEW") // panics
}
```

## Contributing

You can submit a [new issue](https://github.com/nasermirzaei89/env/issues/new) in github [issues](https://github.com/nasermirzaei89/env/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get merged into the main package.
