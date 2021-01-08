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
	check := form.NewCheckbox("Do you need to use form package?", false)
	check2 := form.NewCheckbox("Is this package interesting?", true)
	checkDependance := form.NewCheckbox("Really??", false)
	checkDependance2 := form.NewCheckbox("Really???", false)
	question := form.NewTextField("Any comment? ")

	check2.AddDependance(checkDependance)

	checkDependance.AddDependance(checkDependance2)

	// Add these fields to the form
	myform.Register(check2)
	for _, v := range check2.Children() {
		myform.Register(v)
	}
	myform.Register(title, check, question)

	// Display your form
	myform.Ask()

	// Handle the answers
	fmt.Printf(`use form package? = %v
is this package interesting? = %v
comment = %v
`, check.Answer(), check2.Answer(), question.Answer())
}
