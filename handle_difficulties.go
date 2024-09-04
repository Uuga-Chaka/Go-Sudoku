package main

import (
	"fmt"
	"net/http"
)

type difficulty struct {
	Value  string `json:"value"`
	Option string `json:"option"`
}

func handleDifficulties(w http.ResponseWriter, _ *http.Request) {

	difficulties := []difficulty{
		{"TOO_EASY", "Muy fácil"},
		{"EASY", "Fácil"},
		{"MEDIUM", "Intermedio"},
		{"HARD", "Difícil"},
		{"TOO_HARD", "Muy difícil"},
		{"EXPERT", "Experto"},
	}

	fmt.Println(difficulties)
	respondWithJson(w, http.StatusOK, difficulties)
}
