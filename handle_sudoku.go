package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/go-chi/chi"
)

var difficultyMap = map[string]int{
	"TOO_EASY": 50,
	"EASY":     55,
	"MEDIUM":   60,
	"HARD":     65,
	"TOO_HARD": 70,
	"EXPERT":   75,
}

func handleCreateSudoku(w http.ResponseWriter, r *http.Request) {

	difficulty := chi.URLParam(r, "difficulty")
	fmt.Println("Difficulty level:", difficulty)

	if difficulty == "" {
		difficulty = "TOO_EASY"
	}

	removedCell := difficultyMap[difficulty]

	fmt.Println(r.Body)

	var grid [N][N]int
	num := rand.IntN(8) + 1
	grid[0][0] = num

	solveSudoku(&grid)
	solvedGrid := grid
	removeRandom(removedCell, &grid)
	unsolvedGrid := grid

	type parameters struct {
		Unsolved [N][N]int `json:"unsolved"`
		Solved   [N][N]int `json:"solved"`
		Removed  int       `json:"removed"`
	}

	response := parameters{
		unsolvedGrid,
		solvedGrid,
		removedCell,
	}

	respondWithJson(w, http.StatusOK, response)
}

type Solution struct {
	Solution [N][N]int `json:"solution"`
}

func handleValidateSudoku(w http.ResponseWriter, r *http.Request) {

	var sltn Solution
	err := json.NewDecoder(r.Body).Decode(&sltn)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error")
		return
	}

	if handleSudokuValidation(&sltn.Solution) {
		respondWithJson(w, 200, "Success")
		return
	}

	respondWithJson(w, 200, "Success")
}

func handleSudokuValidation(grid *[N][N]int) bool {

	if validateGroup(grid) {
		fmt.Println(grid)
		fmt.Println("WROOONG")
		return false
	}
	return true
}

func validateHorizontal(grid *[N][N]int) bool {
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			visited := make(map[int]bool)
			num := grid[i][j]
			if visited[num] {
				return false
			}
			visited[num] = true
		}
	}
	return true
}

func validateVertical(grid *[N][N]int) bool {
	for i := 0; i < N; i++ {
		visited := make(map[int]bool)
		for j := 0; j < N; j++ {
			num := grid[i][j]
			if visited[num] {
				return false
			}
			visited[num] = true
		}
	}
	return false
}

func validateGroup(grid *[N][N]int) bool {

	iteration := 0
	maxInteration := 9
	visited := make(map[int]bool)

	for iteration < maxInteration {

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				num := grid[i][j]
				fmt.Println("i", i, "j", j, "Value:", num)
				if visited[num] {
					return false
				}
				visited[num] = true
			}
		}
	}
	return true
}
