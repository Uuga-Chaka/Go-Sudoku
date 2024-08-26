package main

import (
	"log"
	"math/rand/v2"
)

const N = 9

func isValid(grid [N][N]int, row int, column int, num int) bool {
	for i := 0; i < N; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < N; i++ {
		if grid[i][column] == num {
			return false
		}
	}

	startRow, startCol := 3*(row/3), 3*(column/3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(grid *[N][N]int) bool {
	row, col := findEmpty(grid)
	if row == -1 && col == -1 {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValid(*grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func findEmpty(grid *[N][N]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func generateSudokuBoard() [N][N]int {
	var grid [N][N]int

	num := rand.IntN(8) + 1
	grid[0][0] = num

	solveSudoku(&grid)
	removeRandom(34, &grid)

	return grid
}

func removeRandom(quantity int, grid *[N][N]int) {
	if quantity > N*N {
		log.Fatal("You can't remove more elements than the total number of cells")
	}

	removed := 0
	for removed < quantity {
		// Select a random cell
		row := rand.IntN(N)
		col := rand.IntN(N)

		// Check if the cell is already empty
		if grid[row][col] != 0 {
			grid[row][col] = 0
			removed++
		}
	}
}
