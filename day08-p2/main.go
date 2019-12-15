package main

import (
	"fmt"
	color "github.com/fatih/color"
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
	image := make([][25]int, 6)
	digits := make([][10]int, layers)

	for v, _ := range image {
		for h, _ := range image[v] {
			image[v][h] = 2
		}

	}

	for l := 0; l*(vMax*hMax) < len(input); l++ {
		for v := 0; v < vMax; v++ {
			for h := 0; h < hMax; h++ {

				n, _ := strconv.Atoi(string(input[(h)+(v*hMax)+(l*(vMax*hMax))]))
				digits[l][n]++
				if n < 2 && image[v][h] == 2 {
					image[v][h] = n
				}

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

	color.Cyan("Prints text in cyan.")
	c := color.New(color.FgCyan)
	b := color.New(color.FgBlack)
	fmt.Println(leastZero, check, digits[leastZero.layer])
	for _, content := range image {
		for _, n := range content {
			if n == 1 {
				c.Print(n)
				continue
			}
			b.Print(n)
		}
		fmt.Print("\n")

	}
}

type field struct {
	layer int
	count int
}
