package main

import "github.com/samarth-iitb/huh_ordered_multiselect"

func main() {
	var country string
	s := huh.NewSelect[string]().
		Title("Pick a country.").
		Options(
			huh.NewOption("United States", "US"),
			huh.NewOption("Germany", "DE"),
			huh.NewOption("Brazil", "BR"),
			huh.NewOption("Canada", "CA"),
		).
		Value(&country)

	huh.NewForm(huh.NewGroup(s)).Run()
}
