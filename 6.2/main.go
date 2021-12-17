package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fish_per_reproduction_day := make([]int, 9)
	tmp_fish_per_reproduction_day := make([]int, 9)

	lanternfish_str := strings.Split(scanner.Text(), ",")
	for _, lanternfish := range lanternfish_str {
		days_to_reproduction, _ := strconv.Atoi(lanternfish)
		fish_per_reproduction_day[days_to_reproduction]++
	}

	for generation := 1; generation <= 256; generation++ {
		copy(tmp_fish_per_reproduction_day, fish_per_reproduction_day)
		for reproduction_day := 1; reproduction_day <= 8; reproduction_day++ {
			tmp_fish_per_reproduction_day[reproduction_day-1] = fish_per_reproduction_day[reproduction_day]
		}
		tmp_fish_per_reproduction_day[8] = fish_per_reproduction_day[0]
		tmp_fish_per_reproduction_day[6] += fish_per_reproduction_day[0]
		copy(fish_per_reproduction_day, tmp_fish_per_reproduction_day)
	}

	total_number := 0
	for _, number := range fish_per_reproduction_day {
		total_number += number
	}

	fmt.Println()
	fmt.Printf("Number of lanternfish: %d\n", total_number)
}
