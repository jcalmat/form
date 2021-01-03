# Form - Simple form creator library

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jcalmat/form)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Linter](https://github.com/jcalmat/form/workflows/golangci-lint/badge.svg)

Form is a simple library used to create interactive form in your console.

## Usage

```go
package main

import (
	"fmt"

	"github.com/jcalmat/form"
)

func main() {
	myform := form.NewForm()

	title := form.NewLabel("My form")
	check := form.NewCheckbox("Do you need to use form package?", false)
	check2 := form.NewCheckbox("Is this package interesting?", true)
	question := form.NewTextField("Any comment? ")

	myform.Add(title, check, check2, question)

	myform.Ask()

	fmt.Printf(`use form package? = %v
is this package interesting? = %v
comment = %v
`, check.Answer(), check2.Answer(), question.Answer())
}

```

This code will produce the following output

![example](./demo.gif)