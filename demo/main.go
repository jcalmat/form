package main

import (
	"fmt"

	"github.com/jcalmat/form"
)

func main() {
	myform := form.NewForm()

	pasta := form.NewCheckbox("Should I buy pasta?", false)
	beer := form.NewCheckbox("Should I buy beer?", false)
	question := form.NewTextField("Tell me what ya think? ")

	myform.Add(pasta, beer, question)

	myform.Ask()

	fmt.Printf("%v\n%v\n%v\n", pasta.Answer(), beer.Answer(), question.Answer())
}
