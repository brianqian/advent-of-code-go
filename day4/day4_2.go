package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day4_2() {

	fmt.Println("Day 4 -- Pt 2")
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	count := 0
	for s.Scan() {
		pair := strings.Split(s.Text(), ",")
		set1Start, set1End := getStartEnd(pair[0])
		set2Start, set2End := getStartEnd(pair[1])

		if set1End >= set2End || set2Start <= set1Start {
			set1End, set2Start = set2End, set1Start
		}

		if set1End >= set2Start {
			count += 1
		}

	}

	fmt.Println(count)
}
