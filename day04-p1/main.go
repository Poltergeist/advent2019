package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	start := 183564
	end := 657474
	log.WithFields(log.Fields{"start": start, "end": end}).Info()
	valid := 0

	for count := start; count <= end; count++ {

		if checkNumber(count) {
			valid++
		}
	}

	log.Infof("There are %d valid numbers.", valid)

}

func checkNumber(start int) bool {
	last := 0
	two := false
	for _, digit := range fmt.Sprintf("%d", start) {
		i := normalize(digit)
		if last > i {
			return false
		}
		if last == i {
			two = true
		}

		last = i
	}
	return two
}

func normalize(digit rune) int {
	digit = digit - 48
	return int(digit)
}
