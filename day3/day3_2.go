package main

import (
	"bufio"
	"fmt"
	"os"
)

func day3_2() {

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		el1 := makeMap(scanner.Text())
		scanner.Scan()
		el2 := makeMap(scanner.Text())
		scanner.Scan()
		el3 := makeMap(scanner.Text())

		for char := range el1 {
			if el2[char] && el3[char] {
				sum += valMap[string(char)]
			}
		}

	}
	fmt.Println(sum)
}

func makeMap(line string) map[rune]bool {
	obj := make(map[rune]bool)
	for _, char := range line {
		obj[char] = true
	}
	return obj
}
