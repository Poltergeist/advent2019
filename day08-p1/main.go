package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	input := string(dat)
	vMax := 6
	hMax := 25
	layers := len(input) / (vMax * hMax)
	digits := make([][10]int, layers)

	for l := 0; l*(vMax*hMax) < len(input); l++ {
		for v := 0; v < vMax; v++ {
			for h := 0; h < hMax; h++ {

				n, _ := strconv.Atoi(string(input[(h)+(v*hMax)+(l*(vMax*hMax))]))
				digits[l][n]++
				// if l == 97 {

				// 	fmt.Print(n)
				// }

			}
			// if l == 97 {
			// 	fmt.Print("\n")
			// }
		}

	}

	leastZero := field{count: int(^uint(0) >> 1)}

	for l, values := range digits {
		if values[0] < leastZero.count {
			leastZero = field{layer: l, count: values[0]}
		}
	}

	check := digits[leastZero.layer][1] * digits[leastZero.layer][2]

	fmt.Println(leastZero, check, digits[leastZero.layer])
}

type field struct {
	layer int
	count int
}
