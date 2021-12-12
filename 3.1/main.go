package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func apply_bit_count(ones_count []int, number string) {
	for index, bit := range number {
		if bit == '1' {
			ones_count[index] += 1
		}
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	first_line := scanner.Text()

	ones_count := make([]int, len(first_line))
	number_count := 1

	apply_bit_count(ones_count, first_line)

	for scanner.Scan() {
		number_count += 1
		apply_bit_count(ones_count, scanner.Text())
	}

	gamma_str := ""
	epsilon_str := ""

	for _, count := range ones_count {
		if count > number_count/2 {
			gamma_str += "1"
			epsilon_str += "0"
		} else {
			gamma_str += "0"
			epsilon_str += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gamma_str, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilon_str, 2, 64)

	fmt.Printf("Gamma rate: %d\n", gamma)
	fmt.Printf("Epsilon rate: %d\n", epsilon)
	fmt.Printf("Answer: %d\n", gamma*epsilon)
}
