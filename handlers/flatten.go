package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func FlattenMatrix(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	grid, err := getGridFromFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = isValidMatrix(grid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Print each row
	response := strings.Join(grid[0], ",")
	for _, row := range grid[1:] {
		response = fmt.Sprintf("%s,%s", response, strings.Join(row, ","))
	}
	response = response + "\n"
	fmt.Fprint(w, response)
}
