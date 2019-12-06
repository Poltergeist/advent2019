package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type OrbitObject struct {
	orbits *OrbitObject
	name   string
}

func main() {
	orbitList := make([]string, 0)
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		orbitList = append(orbitList, s.Text())
	}
	orbitMap := CreateOrbitMap(orbitList)

	var you, san *OrbitObject
	steps := make([]string, 0)
	for _, object := range orbitMap {
		if object.name == "SAN" {
			san = object
			if len(steps) == 0 {
				steps = countStepsToStart(san, steps)
			} else {
				fmt.Println(countTransfers(san, steps))
			}
		}
		if object.name == "YOU" {
			you = object
			if len(steps) == 0 {
				steps = countStepsToStart(you, steps)
			} else {
				fmt.Println(countTransfers(you, steps))
			}
		}
	}
}

func countTransfers(point *OrbitObject, steps []string) int {
	stepsString := strings.Join(steps, "")
	index := 0
	c := point.orbits
	for c != nil {
		if strings.Contains(stepsString, c.name) {
			for step, object := range steps {
				if object == c.name {

					return step + index
				}
			}
		}
		c = c.orbits
		index++
	}

	return index

}

func countStepsToStart(san *OrbitObject, steps []string) []string {
	c := san.orbits
	for c != nil {
		steps = append(steps, c.name)
		c = c.orbits
	}
	return steps
}

func CreateOrbitMap(orbitList []string) []*OrbitObject {
	orbitMap := make(map[string]*OrbitObject)

	for _, x := range orbitList {
		list := strings.Split(x, ")")
		object := list[1]
		orbits := list[0]

		if orbitMap[orbits] == nil {
			orbitMap[orbits] = &OrbitObject{name: orbits}
		}
		y := orbitMap[orbits]
		if orbitMap[object] != nil {
			orbitMap[object].orbits = y
			continue
		}
		orbitMap[object] = &OrbitObject{name: object, orbits: y}
	}

	r := make([]*OrbitObject, 0)

	for _, orbit := range orbitMap {
		r = append(r, orbit)
	}

	return r

}

func printOrbitMap(orbitMap []*OrbitObject) {
	fmt.Println("printing")
	fmt.Println(orbitMap)
	fmt.Println("done")
	for _, object := range orbitMap {
		fmt.Println(object, object.orbits)
	}
	return
	for _, object := range orbitMap {
		x := object
		for {

			fmt.Print(x.name)
			if x.orbits == nil {
				fmt.Print(x)
				fmt.Print("\n")
				break
			}
			fmt.Print("->")
			x = x.orbits
		}
	}
}
func CalculateOrbitCountCheckSum(orbitMap []*OrbitObject) int {
	checkSum := 0

	for _, object := range orbitMap {
		x := object
		for {
			if x == nil || x.orbits == nil {
				break
			}
			checkSum++
			x = x.orbits
		}
	}

	return checkSum
}
