package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		input := []int{}
		for _, number := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(number)
			input = append(input, n)
		}

		for p1 := 0; p1 < 100; p1++ {
			for p2 := 0; p2 < 100; p2++ {
				output := make([]int, len(input))
				copy(output, input)
				output[1] = p1
				output[2] = p2
				if process(output) == 19690720 {
					log.Info(100*p1 + p2)
					return
				}
			}
		}
	}
}

func process(output []int) int {
	index := 0
	for index < len(output) {
		switch output[index] {
		case 1:
			a := output[output[index+1]]
			b := output[output[index+2]]
			output[output[index+3]] = a + b

			index += 4
		case 2:
			a := output[output[index+1]]
			b := output[output[index+2]]
			output[output[index+3]] = a * b

			index += 4
		case 99:
			return output[0]
		default:
			index++
		}
	}
	return 0
}
