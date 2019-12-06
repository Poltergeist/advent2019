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

	// printOrbitMap(orbitMap)
	fmt.Println(CalculateOrbitCountCheckSum(orbitMap))
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
