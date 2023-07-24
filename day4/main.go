package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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

		if set1Start <= set2Start && set1End >= set2End {
			count += 1
			continue
		}

		if set2Start <= set1Start && set2End >= set1End {
			count += 1
		}

	}

	fmt.Println(count)

	day4_2()
}

func getStartEnd(line string) (int, int) {
	pair := strings.Split(line, "-")
	start, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(err)
	}

	return start, end
}
