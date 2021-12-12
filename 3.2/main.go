package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	numbers := []string{}
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	oxygen_generator_rating_str := find_rating(numbers, most_common_value)
	co2_scrubber_rating_str := find_rating(numbers, least_common_value)

	oxygen_generator_rating, _ := strconv.ParseInt(oxygen_generator_rating_str, 2, 64)
	co2_scrubber_rating, _ := strconv.ParseInt(co2_scrubber_rating_str, 2, 64)

	fmt.Printf("Oxygen generator rating: %d\n", oxygen_generator_rating)
	fmt.Printf("CO2 scrubber rating: %d\n", co2_scrubber_rating)
	fmt.Printf("Answer: %d\n", oxygen_generator_rating*co2_scrubber_rating)
}

func find_rating(numbers []string, filter_fn func([]string, string, int) string) string {
	filter_path := ""
	matching_number := ""
	for index := 0; index < len(numbers[0]); index++ {
		filter_path += filter_fn(numbers, filter_path, index)
		count := 0
		for _, number := range numbers {
			if number[0:(index+1)] == filter_path {
				count += 1
				matching_number = number
			}
		}
		if count == 1 {
			return matching_number
		}
	}
	return ""
}

func count_values(numbers []string, filter_path string, index int) (int, int) {
	zeroes := 0
	ones := 0
	for _, number := range numbers {
		if number[0:(index)] != filter_path {
			continue
		}
		if number[index] == '1' {
			ones += 1
		} else {
			zeroes += 1
		}
	}
	return zeroes, ones
}

func most_common_value(numbers []string, filter_path string, index int) string {
	zeroes, ones := count_values(numbers, filter_path, index)
	if zeroes > ones {
		return "0"
	}
	if ones > zeroes {
		return "1"
	}
	return "1"
}

func least_common_value(numbers []string, filter_path string, index int) string {
	zeroes, ones := count_values(numbers, filter_path, index)
	if zeroes > ones {
		return "1"
	}
	if ones > zeroes {
		return "0"
	}
	return "0"
}
