package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `....#.....#.#...##..........#.......#......
.....#...####..##...#......#.........#.....
.#.#...#..........#.....#.##.......#...#..#
.#..#...........#..#..#.#.......####.....#.
##..#.................#...#..........##.##.
#..##.#...#.....##.#..#...#..#..#....#....#
##...#.............#.#..........#...#.....#
#.#..##.#.#..#.#...#.....#.#.............#.
...#..##....#........#.....................
##....###..#.#.......#...#..........#..#..#
....#.#....##...###......#......#...#......
.........#.#.....#..#........#..#..##..#...
....##...#..##...#.....##.#..#....#........
............#....######......##......#...#.
#...........##...#.#......#....#....#......
......#.....#.#....#...##.###.....#...#.#..
..#.....##..........#..........#...........
..#.#..#......#......#.....#...##.......##.
.#..#....##......#.............#...........
..##.#.....#.........#....###.........#..#.
...#....#...#.#.......#...#.#.....#........
...####........#...#....#....#........##..#
.#...........#.................#...#...#..#
#................#......#..#...........#..#
..#.#.......#...........#.#......#.........
....#............#.............#.####.#.#..
.....##....#..#...........###........#...#.
.#.....#...#.#...#..#..........#..#.#......
.#.##...#........#..#...##...#...#...#.#.#.
#.......#...#...###..#....#..#...#.........
.....#...##...#.###.#...##..........##.###.
..#.....#.##..#.....#..#.....#....#....#..#
.....#.....#..............####.#.........#.
..#..#.#..#.....#..........#..#....#....#..
#.....#.#......##.....#...#...#.......#.#..
..##.##...........#..........#.............
...#..##....#...##..##......#........#....#
.....#..........##.#.##..#....##..#........
.#...#...#......#..#.##.....#...#.....##...
...##.#....#...........####.#....#.#....#..
...#....#.#..#.........#.......#..#...##...
...##..............#......#................
........................#....##..#........#`

	asteroideMap := convertString(input)
	losCountMap := make(map[string]int)
	for yC, _ := range asteroideMap {
		for xC, x := range asteroideMap[yC] {
			if x {
				losCountMap[fmt.Sprintf("%d,%d", xC, yC)] = calculateLineOfSights(asteroideMap, xC, yC)

			}
		}
	}

	fmt.Println(getHighestLosCount(&losCountMap))

}

func getHighestLosCount(losCountMap *map[string]int) (string, int) {
	type point struct {
		location string
		count    int
	}
	highestCount := point{count: 0}

	for key, value := range *losCountMap {
		if value > highestCount.count {
			highestCount = point{location: key, count: value}
		}
	}
	return highestCount.location, highestCount.count
}

func convertString(s string) [][]bool {
	output := [][]bool{}
	for _, x := range strings.Split(s, "\n") {
		l2 := []bool{}
		for _, y := range strings.Split(x, "") {
			l2 = append(l2, y == "#")
		}
		output = append(output, l2)
	}

	return output
}

func calculateLineOfSights(aMap [][]bool, x1 int, y1 int) int {
	losMap := make(map[string]bool)

	for y2, _ := range aMap {
		for x2, point := range aMap[y2] {
			if x2 == x1 && y2 == y1 || !point {
				continue
			}
			if x2 == x1 || y1 == y2 {
				key := ""
				if x2 == x1 {
					if y2 > y1 {
						key = "v"
					} else {
						key = "-v"
					}
				}
				if y2 == y1 {
					if x2 > x1 {
						key = "h"
					} else {
						key = "-h"
					}
				}
				losMap[key] = true
				continue
			}
			var m, b float64
			m = float64(y1-y2) / float64(x1-x2)
			b = float64(y2) - (m * float64(x2))
			key := ""
			if x2 > x1 {
				key = fmt.Sprintf("%f,%f", m, b)
			} else {
				key = fmt.Sprintf("-%f,%f", m, b)
			}
			losMap[key] = true

		}
	}

	// for key, vlaue := range losMap {
	// 	fmt.Println(key, vlaue)
	// }

	return len(losMap)

}
