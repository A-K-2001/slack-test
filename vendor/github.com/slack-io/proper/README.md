# proper [![Build Status](https://github.com/slack-io/proper/actions/workflows/ci.yml/badge.svg)](https://github.com/slack-io/proper/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/slack-io/proper)](https://goreportcard.com/report/github.com/slack-io/proper) [![GoDoc](https://godoc.org/github.com/slack-io/proper?status.svg)](https://godoc.org/github.com/slack-io/proper) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A `map[string]string` decorator offering a collection of helpful functions to extract the values in different types

## Features

* Retrieve data from a string map, in String, Integer, Float and Boolean types.
* Return a default value in case of missing keys or invalid types

# Examples

```go
package main

import (
	"fmt"
	"github.com/slack-io/proper"
)

func main() {
	parameters := make(map[string]string)
	parameters["boolean"] = "true"
	parameters["float"] = "1.2"
	parameters["integer"] = "11"
	parameters["string"] = "value"

	properties := proper.NewProperties(parameters)
	fmt.Println(properties.BooleanParam("boolean", false)) // true
	fmt.Println(properties.FloatParam("float", 0))         // 1.2
	fmt.Println(properties.IntegerParam("integer", 0))     // 11
	fmt.Println(properties.StringParam("string", ""))      // value
}
```