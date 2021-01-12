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

	// Add these fields to the form
	myform.Register(check2).RegisterChildren(checkDependance, checkDependance2).RegisterChild(form.NewLabel("Amazing!"))
	myform.RegisterMany(title, check, question)

	// Display your form
	myform.Run()

	// Handle the answers
	fmt.Printf(`use form package? = %v
is this package interesting? = %v
really?? = %v
really??? = %v
comment = %v
`, check.Answer(), check2.Answer(), checkDependance.Answer(), checkDependance2.Answer(), question.Answer())
}
