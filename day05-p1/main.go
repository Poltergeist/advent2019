package main

import (
	"bufio"
	"fmt"
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
		output := []int{}
		for _, number := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(number)
			output = append(output, n)
		}

		index := 0
		for index < len(output) {
			operation := output[index] % 100
			p1 := output[index] / 100 % 10
			p2 := output[index] / 1000 % 10
			p3 := output[index] / 10000 % 10
			// log.WithFields(log.Fields{"index": index, "operation": operation, "p1": p1, "p2": p2, "p3": p3}).Info()
			switch operation {
			case 1:
				var a, b int
				if p1 == 0 {
					a = output[output[index+1]]
				} else {
					a = output[index+1]
				}
				if p2 == 0 {
					b = output[output[index+2]]
				} else {
					b = output[index+2]
				}
				if p3 == 0 {
					output[output[index+3]] = a + b
				} else {
					output[index+3] = a + b
				}

				index += 4
			case 2:
				var a, b int
				if p1 == 0 {
					a = output[output[index+1]]
				} else {
					a = output[index+1]
				}
				if p2 == 0 {
					b = output[output[index+2]]
				} else {
					b = output[index+2]
				}
				if p3 == 0 {
					output[output[index+3]] = a * b
				} else {
					output[index+3] = a * b
				}

				index += 4
			case 3:
				var inputString string
				fmt.Scanln(&inputString)
				input, _ := strconv.Atoi(inputString)
				if p1 == 0 {
					output[output[index+1]] = input
				} else {
					output[index+1] = input
				}

				index += 2
			case 4:

				if p1 == 0 {
					fmt.Println(output[output[index+1]])
				} else {

					fmt.Println(output[index+1])
				}
				index += 2
			case 99:
				fmt.Println(output)
				return
			default:
				index++
			}
		}
	}
}
