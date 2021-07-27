package main

import (
	"log"

	"github.com/qba73/manhat"
)

func main() {
	if err := manhat.Cli(); err != nil {
		log.Fatal(err)
	}
}
