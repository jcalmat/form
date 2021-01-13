package main

import (
	"github.com/jcalmat/form"
)

func main() {
	// Instantiate a new form
	myform := form.NewForm()

	// Create some fields for your form
	title := form.NewLabel("My form")
	question0 := form.NewCheckbox("Do you need to use form package?", false)
	question0_1 := form.NewTextField("Why? ")

	// Add these fields to the form
	myform.AddItem(title)

	myform.AddItem(question0).
		AddSubItem(question0_1)

	// Display your form
	myform.Run()

	// Handle the answers
	// fmt.Printf("question0 answer = %v, question0_1 answer = %v", question0.Answer(), question0_1.Answer())
}
