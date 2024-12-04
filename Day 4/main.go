package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 4")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()
	inputString := string(input)
	inputSlice := strings.Split(inputString, "\r\n")

	score := 0

	score += checkHorizontal(inputSlice)
	score += checkVertical(inputSlice)
	score += checkDiagonal(inputSlice)

	fmt.Println("Score :", score)
}

func part2() {
	input := ReadInput()
	inputString := string(input)
	field := strings.Split(inputString, "\r\n")

	score := 0

	for i := 1; i < len(field)-1; i++ {
		for j := 1; j < len(field)-1; j++ {
			if string(field[i][j]) == "A" {
				checkWord1 := string(field[i-1][j-1]) + string(field[i][j]) + string(field[i+1][j+1])
				checkWord2 := string(field[i+1][j-1]) + string(field[i][j]) + string(field[i-1][j+1])
				if (checkWord1 == "MAS" || checkWord1 == "SAM") && (checkWord2 == "MAS" || checkWord2 == "SAM") {
					score++
				}
			}
		}
	}

	fmt.Println("Score :", score)
}

func checkHorizontal(field []string) int {
	score := 0

	for _, row := range field {
		for i := 0; i < len(row)-3; i++ {
			checkWord := string(row[i]) + string(row[i+1]) + string(row[i+2]) + string(row[i+3])
			if checkWord == "XMAS" {
				score++
				//fmt.Println("XMAS")
			} else if checkWord == "SAMX" {
				score++
				//Println("SAMX")
			}
		}
	}

	return score
}

func checkVertical(field []string) int {
	score := 0

	for i := 0; i < len(field)-3; i++ {
		for j := 0; j < len(field[i]); j++ {
			checkWord := string(field[i][j]) + string(field[i+1][j]) + string(field[i+2][j]) + string(field[i+3][j])
			if checkWord == "XMAS" {
				score++
				//fmt.Println("XMAS")
			} else if checkWord == "SAMX" {
				score++
				//Println("SAMX")
			}
		}
	}

	return score
}

func checkDiagonal(field []string) int {
	score := 0

	for i := 2; i < len(field)-1; i++ {
		for j := 2; j < len(field[i])-1; j++ {
			checkWord1 := string(field[i-2][j-2]) + string(field[i-1][j-1]) + string(field[i][j]) + string(field[i+1][j+1])
			checkWord2 := string(field[i+1][j-2]) + string(field[i][j-1]) + string(field[i-1][j]) + string(field[i-2][j+1])
			if checkWord1 == "XMAS" {
				score++
			} else if checkWord1 == "SAMX" {
				score++
			}
			if checkWord2 == "XMAS" {
				score++
			} else if checkWord2 == "SAMX" {
				score++
			}
		}
	}

	return score
}

func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
