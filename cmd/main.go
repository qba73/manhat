package main

import (
	"log"
	"os"

	"github.com/qba73/manhat"
)

func main() {
	if err := manhat.Cli(os.Args[1:], os.Stdout); err != nil {
		log.Fatal(err)
	}
}
