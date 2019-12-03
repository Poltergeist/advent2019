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
	wire := 0
	maps := make(map[int]map[string]int)
	for s.Scan() {
		v, h := 0, 0
		maps[wire] = make(map[string]int)
		instructions := strings.Split(s.Text(), ",")
		steps := 0

		for _, instruction := range instructions {
			direction := instruction[:1]
			distance, _ := strconv.Atoi(instruction[1:])
			//log.WithFields(log.Fields{"direction": direction, "distance": distance}).Info("found")
			switch direction {
			case "R":
				start := h
				for h < start+distance {
					h += 1
					steps += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = steps
				}
			case "L":
				start := h
				for h > start-distance {
					h -= 1
					steps += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = steps
				}
			case "U":
				start := v
				for v < start+distance {
					v += 1
					steps += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = steps
				}
			case "D":
				start := v
				for v > start-distance {
					v -= 1
					steps += 1
					maps[wire][fmt.Sprintf("%d:%d", v, h)] = steps
				}

			}

		}
		wire++
	}
	matches := make(map[string]int)
	for key, value := range maps[0] {
		if maps[1][key] != 0 {
			matches[key] = value + maps[1][key]
		}
	}
	steps := int(^uint(0) >> 1)
	for _, value := range matches {
		if steps > value {
			steps = value
		}
	}
	fmt.Println(steps)
}
