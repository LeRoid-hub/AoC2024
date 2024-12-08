package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Day 6")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	var uniqueTiles []string
	direction := 0 // 0=up, 1=right, 2=down, 3=left
	position := []int{0, 0}

	done := false
	for i, row := range inputSlice {
		for j, column := range row {
			if column == '^' {
				position = []int{j, i}
				rowRunes := []rune(row)
				rowRunes[j] = '.'
				inputSlice[i] = string(rowRunes)
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	done = false
	for done == false {
		// get if the next Cell is valid
		nextCellValid := true
		switch direction {
		case 0:
			if position[1]-1 == -1 {
				done = true
				nextCellValid = false
				break
			}
			if inputSlice[position[1]-1][position[0]] == '#' {
				nextCellValid = false
				break
			}
			position[1] = position[1] - 1
			break
		case 1:
			if position[0]+1 == len(inputSlice[0]) {
				done = true
				nextCellValid = false
				break
			}
			if inputSlice[position[1]][position[0]+1] == '#' {
				nextCellValid = false
				break
			}
			position[0] = position[0] + 1
			break
		case 2:
			if position[1]+1 == len(inputSlice) {
				done = true
				nextCellValid = false
				break
			}
			if inputSlice[position[1]+1][position[0]] == '#' {
				nextCellValid = false
				break
			}
			position[1] = position[1] + 1
			break

		case 3:
			if position[0]-1 == -1 {
				done = true
				nextCellValid = false
				break
			}
			if inputSlice[position[1]][position[0]-1] == '#' {
				nextCellValid = false
				break
			}
			position[0] = position[0] - 1
			break

		}

		if nextCellValid {
			num := strings.Join([]string{string(position[0]), string(position[1])}, "")
			if !slices.Contains(uniqueTiles, num) {
				uniqueTiles = append(uniqueTiles, num)
			}
		} else {
			direction++
			direction = direction % 4
		}

	}

	fmt.Println(len(uniqueTiles) + 1)
}

func part2() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	loops := 0
	direction := 0 // 0=up, 1=right, 2=down, 3=left
	position := []int{0, 0}

	done := false
	for i, row := range inputSlice {
		for j, column := range row {
			if column == '^' {
				position = []int{j, i}
				rowRunes := []rune(row)
				rowRunes[j] = '.'
				inputSlice[i] = string(rowRunes)
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	startingPosition := make([]int, len(position))
	copy(startingPosition, position)
	for i, _ := range inputSlice {
		for j, _ := range inputSlice[0] {
			if j == startingPosition[0] && i == startingPosition[1] {
				continue
			}
			copy(position, startingPosition)

			var uniqueTiles []string
			grid := make([]string, len(inputSlice))
			copy(grid, inputSlice)
			rowRunes := []rune(grid[i])
			rowRunes[j] = '#'
			grid[i] = string(rowRunes)

			done = false
			for done == false {
				// get if the next Cell is valid
				nextCellValid := true
				switch direction {
				case 0:
					if position[1]-1 == -1 {
						done = true
						nextCellValid = false
						break
					}
					if grid[position[1]-1][position[0]] == '#' {
						nextCellValid = false
						break
					}
					position[1] = position[1] - 1
					break
				case 1:
					if position[0]+1 == len(grid[0]) {
						done = true
						nextCellValid = false
						break
					}
					if grid[position[1]][position[0]+1] == '#' {
						nextCellValid = false
						break
					}
					position[0] = position[0] + 1
					break
				case 2:
					if position[1]+1 == len(grid) {
						done = true
						nextCellValid = false
						break
					}
					if grid[position[1]+1][position[0]] == '#' {
						nextCellValid = false
						break
					}
					position[1] = position[1] + 1
					break

				case 3:
					if position[0]-1 == -1 {
						done = true
						nextCellValid = false
						break
					}
					if grid[position[1]][position[0]-1] == '#' {
						nextCellValid = false
						break
					}
					position[0] = position[0] - 1
					break

				}

				if nextCellValid {
					num := fmt.Sprintf("%d %d %d", direction, position[0], position[1])
					if !slices.Contains(uniqueTiles, num) {
						uniqueTiles = append(uniqueTiles, num)
					} else {
						done = true
						loops++
						break
					}
				} else {
					direction++
					direction = direction % 4
				}

			}

		}
	}

	fmt.Print(loops)
}

func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
