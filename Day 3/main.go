package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3")
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}

func part1() {
	input := ReadInput()
	inputstring := string(input)
	result := 0

	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

	finds := r.FindAllString(inputstring, -1)

	for _, found := range finds {
		items := strings.Split(found, ",")
		a := items[0]
		b := items[1]
		numa := a[4:]
		numb := b[:len(b)-1]

		inta, err := strconv.Atoi(numa)
		if err != nil {
			panic(err)
		}
		intb, err := strconv.Atoi(numb)
		if err != nil {
			panic(err)
		}

		result += (inta * intb)

	}
	fmt.Println(result)
}

func part2() {
	input := ReadInput()
	inputstring := string(input)
	result := 0
	condition := true

	r, _ := regexp.Compile("(mul\\([0-9]{1,3},[0-9]{1,3}\\))|(do\\(\\))|(don't\\(\\))")

	finds := r.FindAllString(inputstring, -1)

	for _, found := range finds {
		if found == "do()" {
			condition = true
		} else if found == "don't()" {
			condition = false
		} else if condition == true {
			items := strings.Split(found, ",")
			a := items[0]
			b := items[1]
			numa := a[4:]
			numb := b[:len(b)-1]

			inta, err := strconv.Atoi(numa)
			if err != nil {
				panic(err)
			}
			intb, err := strconv.Atoi(numb)
			if err != nil {
				panic(err)
			}

			result += (inta * intb)
		}

	}
	fmt.Println(result)

}
func ReadInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return file
}
