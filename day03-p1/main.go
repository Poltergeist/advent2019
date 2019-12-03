package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
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
	wire := 0
	maps := make(map[int]map[string]bool)
	for s.Scan() {
		v, h := 0, 0
		maps[wire] = make(map[string]bool)
		instructions := strings.Split(s.Text(), ",")

		for _, instruction := range instructions {
			direction := instruction[:1]
			distance, _ := strconv.Atoi(instruction[1:])
			log.WithFields(log.Fields{"direction": direction, "distance": distance}).Info("found")
			switch direction {
			case "R":
				start := h
				for h < start+distance {
					h += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = true
				}
			case "L":
				start := h
				for h > start-distance {
					h -= 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = true
				}
			case "U":
				start := v
				for v < start+distance {
					v += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = true
				}
			case "D":
				start := v
				for v > start-distance {
					v -= 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = true
				}

			}

		}
		wire++
	}
	matches := make(map[string]bool)
	for key, value := range maps[0] {
		if maps[1][key] == value {
			matches[key] = value
		}
	}
	max := int(^uint(0) >> 1)
	for key, _ := range matches {
		sum := 0
		for _, i := range strings.Split(key, ":") {
			value, _ := strconv.ParseFloat(i, 64)
			sum += int(math.Abs(value))
		}
		if max > sum {
			max = sum
		}
	}
	fmt.Println(max)
}
