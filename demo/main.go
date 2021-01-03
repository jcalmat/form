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
