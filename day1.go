package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func check(err error, id int) {
	if err != nil {
		fmt.Println(id)
		log.Fatal("Error!")
	}
}

func Day1() {
	fmt.Println("Day 1")

	f, err := os.Open("./day1data.txt")
	check(err, 1)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	current := []int{}
	max := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			sum := 0
			for _, val := range current {
				sum += val
			}
			max = int(math.Max(float64(max), float64(sum)))
			current = []int{}
		} else {
			asInt, err := strconv.Atoi(text)
			check(err, 2)
			current = append(current, asInt)
		}
	}

	fmt.Println(max)

}
