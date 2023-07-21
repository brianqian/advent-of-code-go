package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main2() {

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 0

	rps := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	toPoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		villain := c[0]
		hero := c[1]

		result, err := determineChoice(villain, hero, toPoints, rps)
		if err != nil {
			panic(err)
		}

		score += result
	}

	fmt.Println(score)

}

func determineChoice(villain string, outcome string, p map[string]int, r map[string]string) (int, error) {
	if outcome == "Y" {
		return p[r[villain]] + 3, nil
	}

	if outcome == "X" {
		if villain == "A" {
			return p["Z"], nil
		}

		if villain == "B" {
			return p["X"], nil
		}

		if villain == "C" {
			return p["Y"], nil
		}
	}
	if outcome == "Z" {
		if villain == "A" {
			return p["Y"] + 6, nil
		}

		if villain == "B" {
			return p["Z"] + 6, nil
		}

		if villain == "C" {
			return p["X"] + 6, nil
		}
	}

	return -1, errors.New("invalid input")
}
