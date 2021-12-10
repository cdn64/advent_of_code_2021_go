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

	line := 0
	window := [...]int{-1, -1, -1}

	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())

		window = [...]int{window[1], window[2], depth}
		if line > 1 {
			depth := 0
			for _, v := range window {
				depth += v
			}

			if last_depth == -1 {
				fmt.Printf("%d: (N/A - no previous measurement)\n", depth)
			} else if depth > last_depth {
				fmt.Printf("%d: (increased)\n", depth)
				increased_count += 1
			} else if depth < last_depth {
				fmt.Printf("%d: (decreased)\n", depth)
			} else {
				fmt.Printf("%d: (no change)\n", depth)
			}
			last_depth = depth
		}

		line += 1
	}

	fmt.Printf("Increased %d times\n", increased_count)
}
