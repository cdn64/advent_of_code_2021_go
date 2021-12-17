package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type OceanFloor struct {
	vent_count [][]int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	ocean_floor := OceanFloor{}

	for scanner.Scan() {
		coordinates := [4]int{}
		for i, endpoint := range strings.Split(scanner.Text(), " -> ") {
			for j, coordinate := range strings.Split(endpoint, ",") {
				coordinates[(i*2)+j], _ = strconv.Atoi(coordinate)
			}
		}

		if coordinates[0] == coordinates[2] {
			for y := min(coordinates[1], coordinates[3]); y <= max(coordinates[1], coordinates[3]); y++ {
				ocean_floor = ensure_size(ocean_floor, coordinates[0], y)
				ocean_floor.vent_count[coordinates[0]][y]++
			}
		}
		if coordinates[1] == coordinates[3] {
			for x := min(coordinates[0], coordinates[2]); x <= max(coordinates[0], coordinates[2]); x++ {
				ocean_floor = ensure_size(ocean_floor, x, coordinates[1])
				ocean_floor.vent_count[x][coordinates[1]]++
			}
		}
	}

	dangerous_areas := 0
	for _, row := range ocean_floor.vent_count {
		for _, vent_count := range row {
			if vent_count > 1 {
				dangerous_areas += 1
			}
		}
	}

	fmt.Println()
	fmt.Printf("Number of dangerous areas: %d\n", dangerous_areas)
}

func ensure_size(ocean_floor OceanFloor, x int, y int) OceanFloor {
	for x >= len(ocean_floor.vent_count) {
		ocean_floor.vent_count = append(ocean_floor.vent_count, []int{})
	}
	for y >= len(ocean_floor.vent_count[x]) {
		ocean_floor.vent_count[x] = append(ocean_floor.vent_count[x], 0)
	}
	return ocean_floor
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
