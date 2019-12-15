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
	max := 0

	s := bufio.NewScanner(f)
	out := 0
	maxVariant := make([]int, 5)
	variants := generateVariants()
	getOut := func() int {
		// fmt.Println(out)
		return out
	}
	inputPrompt := []func() int{getInputFromAuto, getOut}
	originalOutput := []int{}
	for s.Scan() {
		for _, number := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(number)
			originalOutput = append(originalOutput, n)
		}

	}

	// fmt.Println(variants)
	for _, variant := range variants {
		a = variant
		out = 0
		output := make([]int, len(originalOutput))
		copy(output, originalOutput)
		// fmt.Println("reset Output", output, a)
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
				input := inputPrompt[0]()
				inputPrompt = []func() int{inputPrompt[1], inputPrompt[0]}
				// fmt.Println("input", input)
				if p1 == 0 {
					output[output[index+1]] = input
				} else {
					output[index+1] = input
				}

				index += 2
			case 4:

				if p1 == 0 {
					out = output[output[index+1]]
				} else {
					out = output[index+1]
				}
				// fmt.Println("output", out)
				if out > max {

					max = out
					copy(maxVariant, variant)
				}
				index += 2
			case 5:
				// jump if true p1 != 0
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

				index += 3
				if a != 0 {
					index = b
				}
			case 6:
				// jump if false p1 == 0
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
				index += 3
				if a == 0 {
					index = b
				}
			case 7:
				// less then  p1 < p2 ? p3 = 1 : p3 = 0
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

				set := 0
				if a < b {
					set = 1

				}

				if p3 == 0 {
					output[output[index+3]] = set
				} else {
					output[index+3] = set
				}
				index += 4
			case 8:
				// equal p1 == p2 ? p3 = 1 : p3 = 0
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

				set := 0
				if a == b {
					set = 1
				}

				if p3 == 0 {
					output[output[index+3]] = set
				} else {
					output[index+3] = set
				}
				index += 4
			case 99:
				if len(a) != 0 {
					index = 0
					// fmt.Println("next")
					continue
				}
				log.Info(out, "end")
				index = len(output) + 1
			default:
				index++
			}
		}
	}
	fmt.Println(max, maxVariant, len(variants))

}

var a []int = []int{4, 3, 2, 1, 0}

func generateVariants() [][]int {
	start := []int{0, 1, 2, 3, 4}
	// start := []int{4, 3, 2, 1, 0}
	var returnValue [][]int
	// returnValue = append(returnValue, start)
	// return returnValue
	Perm(start, func(a []int) {
		x := make([]int, len(a))
		copy(x, a)
		returnValue = append(returnValue, x)
	})

	return returnValue

}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
func getInputFromAuto() int {
	var x int
	if len(a) == 0 {
		fmt.Println("Array empty")
		return 0
	}
	x, a = a[0], a[1:]
	fmt.Println(x, a)
	return x
}

// func getInputFromStdin() int {
// 	var inputString string
// 	fmt.Scanln(&inputString)
// 	input, _ := strconv.Atoi(inputString)
// 	return input
// }
