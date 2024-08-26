package main

import (
	"math/rand/v2"
	"net/http"
)

func handleCreateSudoku(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Unsolved [N][N]int `json:"unsolved"`
		Solved   [N][N]int `json:"solved"`
	}

	var grid [N][N]int
	num := rand.IntN(8) + 1
	grid[0][0] = num

	solveSudoku(&grid)
	solvedGrid := grid
	removeRandom(37, &grid)
	unsolvedGrid := grid

	response := parameters{
		Unsolved: unsolvedGrid,
		Solved:   solvedGrid,
	}

	respondWithJson(w, 200, response)
}
