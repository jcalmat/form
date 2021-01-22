# Form - Simple form creator library

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jcalmat/form)
[![License: MIT](https://img.shields.io/github/license/jcalmat/form)](https://opensource.org/licenses/MIT)
![Linter](https://github.com/jcalmat/form/workflows/golangci-lint/badge.svg)

Form is a simple library used to create interactive form in your console.

## Usage

Code available [here](./example/main.go)
```go
package main

import (
	"fmt"

	"github.com/jcalmat/form"
)

func main() {
	// Instantiate a new form
	myform := form.NewForm()

	// Create some fields for your form
	title := form.NewLabel("My form")
	question0 := form.NewCheckbox("Do you need to use form package?", false)
	question0_1 := form.NewTextField("Why? ")
	question0_2 := form.NewCheckbox("Really? ", false)

	// Add these fields to the form
	myform.AddItem(title)

	// Add items dependant on one another
	myform.AddItem(question0).
		AddItem(question0_1).
		AddItem(question0_2).
		AddItem(form.NewLabel("Amazing!"))

	// Display your form
	myform.Run()

	// Handle the answers
	fmt.Printf("question0 answer = %v, question0_1 answer = %v, question0_2 answer = %v\n", question0.Answer(), question0_1.Answer(), question0_2.Answer())
}
```

This code will produce the following output

![example](./example/demo.gif)