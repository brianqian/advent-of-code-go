package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main_2() {
	fmt.Println("Day 1")

	f, err := os.Open("./day1data.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	current := 0
	nums := []int{0, 0, 0}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			nums = append(nums, current)
			current = 0
			sort.Ints(nums)
			nums = nums[1:]
		} else {
			c, err := strconv.Atoi(text)
			check(err)
			current += c
		}
	}

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[0] + nums[1] + nums[2])

}
