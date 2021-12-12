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

func find_rating(numbers []string, filter_fn func([]string, int) byte) string {
	rating_str := ""
	filtering_numbers := append([]string(nil), numbers...)
	for index := 0; index < len(numbers); index++ {
		filter_value := filter_fn(filtering_numbers, index)
		filtered_numbers := []string{}
		for _, number := range filtering_numbers {
			if number[index] == filter_value {
				filtered_numbers = append(filtered_numbers, number)
			}
		}
		if len(filtered_numbers) == 1 {
			rating_str = filtered_numbers[0]
			break
		}
		filtering_numbers = filtered_numbers
	}
	return rating_str
}

func count_values(numbers []string, index int) (int, int) {
	zeroes := 0
	ones := 0
	for _, number := range numbers {
		if number[index] == '1' {
			ones += 1
		} else {
			zeroes += 1
		}
	}
	return zeroes, ones
}

func most_common_value(numbers []string, index int) byte {
	zeroes, ones := count_values(numbers, index)
	if zeroes > ones {
		return '0'
	}
	if ones > zeroes {
		return '1'
	}
	return '1'
}

func least_common_value(numbers []string, index int) byte {
	zeroes, ones := count_values(numbers, index)
	if zeroes > ones {
		return '1'
	}
	if ones > zeroes {
		return '0'
	}
	return '0'
}
