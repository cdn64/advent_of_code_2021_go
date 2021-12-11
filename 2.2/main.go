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

	aim := 0
	depth := 0
	position := 0

	for scanner.Scan() {
		command := scanner.Text()

		if strings.Contains(command, "forward") {
			movement, _ := strconv.Atoi(command[8:])
			position += movement
			depth += movement * aim
		} else if strings.Contains(command, "down") {
			movement, _ := strconv.Atoi(command[5:])
			aim += movement
		} else if strings.Contains(command, "up") {
			movement, _ := strconv.Atoi(command[3:])
			aim -= movement
		}
	}

	fmt.Printf("Final depth: %d\n", depth)
	fmt.Printf("Final position: %d\n", position)
	fmt.Printf("Answer: %d\n", depth*position)
}
