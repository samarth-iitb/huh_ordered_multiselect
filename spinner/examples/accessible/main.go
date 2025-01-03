package main

import (
	"context"
	"log"
	"time"

	"github.com/samarth-iitb/huh_ordered_multiselect/spinner"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()

	err := spinner.New().
		Context(ctx).
		Accessible(true).
		Run()
	if err != nil {
		log.Fatalln(err)
	}
}
