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

	depth := 0
	position := 0

	for scanner.Scan() {
		command := scanner.Text()

		if strings.Contains(command, "forward") {
			movement, _ := strconv.Atoi(command[8:])
			position += movement
		} else if strings.Contains(command, "down") {
			movement, _ := strconv.Atoi(command[5:])
			depth += movement
		} else if strings.Contains(command, "up") {
			movement, _ := strconv.Atoi(command[3:])
			depth -= movement
		}
	}

	fmt.Printf("Final depth: %d\n", depth)
	fmt.Printf("Final position: %d\n", position)
	fmt.Printf("Answer: %d\n", depth*position)
}
