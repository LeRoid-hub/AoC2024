package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	dict := make([][]int, 100)
	var updates [][]int
	count := 0

	for i, row := range inputSlice {
		if row == "" {
			for j := i + 1; j < len(inputSlice); j++ {
				temp := strings.Split(string(inputSlice[j]), ",")
				tempInt := make([]int, len(temp))
				for i, s := range temp {
					tempInt[i], _ = strconv.Atoi(s)
				}
				updates = append(updates, tempInt)
			}
			break
		} else {

			temp := strings.Split(row, "|")
			tempIndex, _ := strconv.Atoi(temp[0])
			tempAfter, _ := strconv.Atoi(temp[1])
			dict[tempIndex] = append(dict[tempIndex], tempAfter)
		}
	}

	for _, row := range updates {
		valid := true

		for i := 1; i < len(row); i++ {
			rules := dict[row[i]]
			if slices.Contains(rules, row[i-1]) {
				valid = false
				break
			}
		}

		if valid {
			count += row[len(row)/2]
		}
	}

	fmt.Println(count)
}

func part2() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	dict := make([][]int, 100)
	var updates [][]int
	count := 0

	for i, row := range inputSlice {
		if row == "" {
			for j := i + 1; j < len(inputSlice); j++ {
				temp := strings.Split(string(inputSlice[j]), ",")
				tempInt := make([]int, len(temp))
				for i, s := range temp {
					tempInt[i], _ = strconv.Atoi(s)
				}
				updates = append(updates, tempInt)
			}
			break
		} else {

			temp := strings.Split(row, "|")
			tempIndex, _ := strconv.Atoi(temp[0])
			tempAfter, _ := strconv.Atoi(temp[1])
			dict[tempIndex] = append(dict[tempIndex], tempAfter)
		}
	}

	for i, row := range updates {
		valid := true
		for j := 1; j < len(row); j++ {
			rules := dict[row[j]]
			if slices.Contains(rules, row[j-1]) {
				valid = false
				tempRow := make([]int, len(row))
				tempSorted := make([]int, len(tempRow))
				copy(tempRow, row)

				tempSorted[0] = tempRow[0]

				for h, temp := range tempRow {
					newRules := dict[temp]
					currentBest := 0

					if h == 0 {
						continue
					}

					for k := 0; k < len(tempSorted); k++ {
						oldRules := dict[tempSorted[k]]

						if slices.Contains(oldRules, temp) {
							currentBest = k + 1
							continue
						}

						if slices.Contains(newRules, tempSorted[k]) {
							currentBest = k
							break
						}

					}

					tempfinal := make([]int, len(tempSorted))
					for k := 0; k < len(tempfinal); k++ {
						if k < currentBest {
							tempfinal[k] = tempSorted[k]
						} else if k == currentBest {
							tempfinal[k] = temp
						} else {
							tempfinal[k] = tempSorted[k-1]
						}
					}
					tempSorted = tempfinal

				}

				updates[i] = tempSorted
			}
		}

		if !valid {
			counter := updates[i]
			count += counter[len(row)/2]
		}
	}

	fmt.Println(count)
}

func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
