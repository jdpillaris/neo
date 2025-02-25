package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func TransposeMatrix(w http.ResponseWriter, r *http.Request) {
	grid, err := getGridFromFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = getIntMatrix(grid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transposedMatrix := transpose(grid)
	// Print each row
	var response string
	for _, row := range transposedMatrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func transpose(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]string, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}
