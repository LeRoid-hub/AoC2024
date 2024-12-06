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
}

func part1() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	var uniqueTiles []int
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
		fmt.Println(position, direction)
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
			if position[0]+1 == len(inputSlice[0])+1 {
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
			if position[1]+1 == len(inputSlice)+1 {
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

			var padding int = 1
			for padding < position[1] {
				padding *= 10
			}
			num := position[0]*padding + position[1]
			if !slices.Contains(uniqueTiles, num) {
				uniqueTiles = append(uniqueTiles, num)
			}
		} else {
			fmt.Println("TURN AROUND")
			direction++
			direction = direction % 4
			fmt.Println(direction)
		}

	}

	fmt.Println(direction, position)

	fmt.Println(inputSlice)
	fmt.Println(uniqueTiles)
	fmt.Println(len(uniqueTiles))
}

func ReadInput() []byte {
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	return file
}
