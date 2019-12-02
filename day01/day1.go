package main

import (
	"fmt"
	"log"
	m "math"
)

func main() {
	var n float64
	input := "1969"
	_, err := fmt.Sscanf(input, "%f", &n)
	if err != nil {
		log.Fatalf("could not read %s: %v", input, err)
	}
	fmt.Println(m.Floor(n/3) - 2)
}
