package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MARKED_CELL = "x"

type BingoBoardRow struct {
	cells []string
}
type BingoBoard struct {
	rows []BingoBoardRow
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers_to_draw := strings.Split(scanner.Text(), ",")
	fmt.Printf("%d numbers to draw\n", len(numbers_to_draw))

	boards := []BingoBoard{}
	rows := []BingoBoardRow{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		row := strings.Fields(line)
		rows = append(rows, BingoBoardRow{row})
		if len(rows) == 5 {
			boards = append(boards, BingoBoard{rows})
			rows = []BingoBoardRow{}
		}
	}

	fmt.Printf("%d boards parsed\n", len(boards))

	last_number := 0
	winning_boards := make(map[int]bool)
	var last_winning_board BingoBoard

	for _, number := range numbers_to_draw {
		for index, board := range boards {
			mark_number(board, number)
			if has_won(board) {
				_, ok := winning_boards[index]
				if !ok {
					winning_boards[index] = true
					last_winning_board = board
				}
			}
		}
		if len(winning_boards) == len(boards) {
			last_number, _ = strconv.Atoi(number)
			break
		}
	}

	fmt.Println("Losing board:")
	unmarked_numbers_sum := 0
	for _, row := range last_winning_board.rows {
		for _, cell := range row.cells {
			fmt.Printf("%-5v", cell)
			if cell != MARKED_CELL {
				cell_as_int, _ := strconv.Atoi(cell)
				unmarked_numbers_sum += cell_as_int
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("Sum of all unmarked numbers: %d\n", unmarked_numbers_sum)
	fmt.Printf("Last number: %d\n", last_number)
	fmt.Printf("Answer: %d\n", unmarked_numbers_sum*last_number)
}

func mark_number(board BingoBoard, number string) {
	for row, row_content := range board.rows {
		for cell, _ := range row_content.cells {
			if board.rows[row].cells[cell] == number {
				board.rows[row].cells[cell] = MARKED_CELL
				return
			}
		}
	}
}

func has_won(board BingoBoard) bool {
	board_size := len(board.rows)
	for row := 0; row < board_size; row++ {
		has_unmarked_cells := false
		for column := 0; column < board_size; column++ {
			if board.rows[row].cells[column] != MARKED_CELL {
				has_unmarked_cells = true
				break
			}
		}
		if !has_unmarked_cells {
			return true
		}
	}
	for column := 0; column < board_size; column++ {
		has_unmarked_cells := false
		for row := 0; row < board_size; row++ {
			if board.rows[row].cells[column] != MARKED_CELL {
				has_unmarked_cells = true
				break
			}
		}
		if !has_unmarked_cells {
			return true
		}
	}
	return false
}
