package main

import (
	"log"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
