package main

import "github.com/samarth-iitb/huh_ordered_multiselect"

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("First"),
			huh.NewInput().Title("Second"),
			huh.NewInput().Title("Third"),
			huh.NewInput().Title("Fourth"),
			huh.NewInput().Title("Fifth"),
			huh.NewInput().Title("Sixth"),
			huh.NewInput().Title("Seventh"),
			huh.NewInput().Title("Eigth"),
			huh.NewInput().Title("Nineth"),
			huh.NewInput().Title("Tenth"),
		),
	).WithHeight(5)
	form.Run()
}
