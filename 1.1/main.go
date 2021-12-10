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

	last_depth := -1
	increased_count := 0

	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())

		if last_depth == -1 {
			fmt.Printf("%d: (N/A - no previous measurement)\n", depth)
		} else if depth > last_depth {
			fmt.Printf("%d: (increased)\n", depth)
			increased_count += 1
		} else {
			fmt.Printf("%d: (decreased)\n", depth)
		}

		last_depth = depth
	}

	fmt.Printf("Increased %d times\n", increased_count)
}
