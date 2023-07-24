package main

import (
	"bufio"
	"fmt"
	"os"
)

var valMap = make(map[string]int)

func main() {
	fmt.Printf("Day 3 \n")
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		val := getValueOfRepeat((scanner.Text()))
		sum += val
	}

	fmt.Println(sum)

	day3_2()

}

func getValueOfRepeat(line string) int {
	midpoint := len(line) / 2
	cat1 := line[:midpoint]
	cat2 := line[midpoint:]

	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, l := range alpha {
		valMap[string(l)] = i + 1
	}

	m := make(map[rune]bool)

	for _, r := range cat1 {
		m[r] = true
	}

	for _, r := range cat2 {
		if m[r] {
			return valMap[string(r)]
		}
	}
	panic("Error")
}
