package main

import (
	"bufio"
	"fmt"
	"log"
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
		output := []int{}
		for _, number := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(number)
			output = append(output, n)
		}

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
				fmt.Println(output[0])
				return
			default:
				index++
			}
		}
	}
}
