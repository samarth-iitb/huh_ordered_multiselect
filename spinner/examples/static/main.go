package main

import (
	"fmt"

	"github.com/samarth-iitb/huh_ordered_multiselect/spinner"
)

func main() {
	_ = spinner.New().Title("Loading").Accessible(true).Run()
	fmt.Println("Done!")
}
