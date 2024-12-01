package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()

	inputstring := string(input)

	inputrow := strings.Split(inputstring, "\r")

	var leftgroup []int
	var rightgroup []int
	for _, row := range inputrow {
		splitstr := strings.Fields(row)

		num, err := strconv.Atoi(splitstr[0])
		if err != nil {
			panic(err)
		}
		leftgroup = append(leftgroup, num)

		num, err = strconv.Atoi(splitstr[1])
		if err != nil {
			panic(err)
		}
		rightgroup = append(rightgroup, num)
	}

	sort.Slice(leftgroup, func(i, j int) bool {
		return leftgroup[i] < leftgroup[j]
	})

	sort.Slice(rightgroup, func(i, j int) bool {
		return rightgroup[i] < rightgroup[j]
	})

	var distance int

	for i := 0; i < len(leftgroup); i++ {
		if leftgroup[i] > rightgroup[i] {
			distance += (leftgroup[i] - rightgroup[i])
		} else if rightgroup[i] > leftgroup[i] {
			distance += (rightgroup[i] - leftgroup[i])
		}
	}

	fmt.Println(distance)
}

func part2() {
	input := ReadInput()

	inputstring := string(input)

	inputrow := strings.Split(inputstring, "\r")

	var leftgroup []int
	var rightgroup []int
	for _, row := range inputrow {
		splitstr := strings.Fields(row)

		num, err := strconv.Atoi(splitstr[0])
		if err != nil {
			panic(err)
		}
		leftgroup = append(leftgroup, num)

		num, err = strconv.Atoi(splitstr[1])
		if err != nil {
			panic(err)
		}
		rightgroup = append(rightgroup, num)
	}

	sort.Slice(leftgroup, func(i, j int) bool {
		return leftgroup[i] < leftgroup[j]
	})

	sort.Slice(rightgroup, func(i, j int) bool {
		return rightgroup[i] < rightgroup[j]
	})

	m := make(map[int]int)
	score := 0

	for _, num := range rightgroup {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	for _, num := range leftgroup {
		if _, ok := m[num]; ok {
			if m[num] > 0 {
				score += (num * m[num])
			}
		}
	}

	fmt.Println(score)
}
func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
