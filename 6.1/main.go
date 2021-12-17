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

	lanternfish_str := strings.Split(scanner.Text(), ",")
	lanternfish := make([]int, len(lanternfish_str))
	for i, l := range lanternfish_str {
		lanternfish[i], _ = strconv.Atoi(l)
	}

	for generation := 1; generation <= 80; generation++ {
		generation_length := len(lanternfish)
		for i := 0; i < generation_length; i++ {
			lanternfish[i]--
			if lanternfish[i] < 0 {
				lanternfish[i] = 6
				lanternfish = append(lanternfish, 8)
			}
		}
	}

	fmt.Println()
	fmt.Printf("Number of lanternfish: %d\n", len(lanternfish))
}
