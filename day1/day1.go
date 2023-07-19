package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error!")
	}
}

func main() {
	fmt.Println("Day 1")

	f, err := os.Open("./day1data.txt")
	check(err)
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
			check(err)
			current = append(current, asInt)
		}
	}

	fmt.Println(max)
	main_2()
}
