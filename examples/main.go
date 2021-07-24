package main

import (
	"fmt"
	"log"

	"github.com/qba73/manhat"
)

func main() {
	// define a point value
	point := 23

	// calculate distance from the point to the center
	distance, err := manhat.CalculateDistance(point)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got distance: ", distance)
}
