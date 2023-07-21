package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type RpsScore struct {
	X int8
	Y int8
	Z int8
}

func main() {

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 0

	rps := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		villain := c[0]
		hero := c[1]

		result, err := determineWinner(villain, hero)
		if err != nil {
			panic(err)
		}

		score += result + rps[hero]
	}

	fmt.Println(score)

	fmt.Println("Day 2")
	main2()

}

func determineWinner(villain string, hero string) (int, error) {
	var h string
	if hero == "X" {
		h = "A"
	} else if hero == "Y" {
		h = "B"
	} else if hero == "Z" {
		h = "C"
	}

	if villain == h {
		return 3, nil
	}

	if villain == "A" {
		if h == "B" {
			return 6, nil
		}
		return 0, nil
	}

	if villain == "B" {
		if h == "C" {
			return 6, nil
		}
		return 0, nil
	}

	if villain == "C" {
		if h == "A" {
			return 6, nil
		}
		return 0, nil
	}

	return -1, errors.New("invalid input")
}
