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

func checkNumber(count int) bool {
	last := 0
	double := make(map[int]string)
	for _, digit := range fmt.Sprintf("%d", count) {
		i := Normalize(digit)
		if last > i {
			return false
		}
		if last == i {
			switch double[i] {
			case "":
				double[i] = "yes"
			case "yes":

				double[i] = "no"
			}
		}

		last = i
	}
	for _, value := range double {
		if value == "yes" {
			return true
		}
	}
	return false
}

func Normalize(digit rune) int {
	digit = digit - 48
	return int(digit)
}
