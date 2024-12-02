package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()
	var safe int

	inputstring := string(input)

	reports := strings.Split(inputstring, "\r\n")

	for _, report := range reports {
		levels := strings.Split(report, " ")
		increase := 0 // 0 = init, 1 = decrease, 2 = increase
		issafe := true

		for i := 0; i < len(levels)-1; i++ {
			level1, err1 := strconv.Atoi(levels[i])
			if err1 != nil {
				panic(err1)
			}
			level2, err2 := strconv.Atoi(levels[i+1])
			if err2 != nil {
				panic(err2)
			}

			if increase == 0 {
				if level1 > level2 {
					increase = 1
				} else if level1 < level2 {
					increase = 2
				}
			}

			if !isSafe(level1, level2, increase) {
				issafe = false
				break
			}

		}

		if issafe == true {
			safe++
		}

	}

	fmt.Println(safe)
}

func isSafe(base int, compare int, increase int) bool {
	distance := float64(base - compare)
	distance = math.Abs(distance)

	if distance > 3 {
		return false
	} else if distance < 1 {
		return false
	}

	if increase == 1 {
		if base > compare {
			return true
		}
	} else if increase == 2 {
		if base < compare {
			return true
		}
	}

	return false
}

func part2() {
	input := ReadInput()
	var safe int

	inputstring := string(input)

	reports := strings.Split(inputstring, "\r\n")

	var shits []string

	for _, report := range reports {
		levels := strings.Split(report, " ")
		increase := 0 // 0 = init, 1 = decrease, 2 = increase
		issafe := true

		for i := 0; i < len(levels)-1; i++ {
			level1, err1 := strconv.Atoi(levels[i])
			if err1 != nil {
				panic(err1)
			}
			level2, err2 := strconv.Atoi(levels[i+1])
			if err2 != nil {
				panic(err2)
			}

			if increase == 0 {
				if level1 > level2 {
					increase = 1
				} else if level1 < level2 {
					increase = 2
				}
			}

			if !isSafe(level1, level2, increase) {
				issafe = false
				shits = append(shits, report)
				break
			}

		}

		if issafe == true {
			safe++
		}

	}

	for _, shit := range shits {
		shitlevel := strings.Split(shit, " ")

		for i := 0; i < len(shitlevel); i++ {
			shitsafe := true
			shitcrease := 0
			holdcopyshit := make([]string, len(shitlevel))
			copy(holdcopyshit, shitlevel)
			copyshit := remove(holdcopyshit, i)

			for i := 0; i < len(copyshit)-1; i++ {
				level1, err1 := strconv.Atoi(string(copyshit[i]))
				if err1 != nil {
					panic(err1)
				}
				level2, err2 := strconv.Atoi(string(copyshit[i+1]))
				if err2 != nil {
					panic(err2)
				}

				if shitcrease == 0 {
					if level1 > level2 {
						shitcrease = 1
					} else if level1 < level2 {
						shitcrease = 2
					}
				}

				if !isSafe(level1, level2, shitcrease) {
					shitsafe = false
					break
				}

			}

			if shitsafe == true {
				safe++
				break
			}
		}

	}

	fmt.Println(safe)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
