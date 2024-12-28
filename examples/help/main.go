package main

import "github.com/samarth-iitb/huh_ordered_multiselect"

func main() {
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Dynamic Help"),
			huh.NewInput().Title("Dynamic Help"),
			huh.NewInput().Title("Dynamic Help"),
		),
	)
	f.Run()
}
