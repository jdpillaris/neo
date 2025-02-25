package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func EchoMatrix(w http.ResponseWriter, r *http.Request) {
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
	var response string
	for _, row := range grid {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

// Parses the file from the request into a matrix
func getGridFromFile(r *http.Request) ([][]string, error) {
	// Parse file upload
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read CSV data
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func isValidMatrix(matrix [][]string) error {
	matrixLength := len(matrix)
	if matrixLength == 0 {
		return fmt.Errorf("file is empty")
	}

	for i, row := range matrix {
		if len(row) != matrixLength {
			return fmt.Errorf("number of rows not equal to size of row %d", i)
		}
		for j, cell := range row {
			if _, err := strconv.Atoi(cell); err != nil {
				return fmt.Errorf("invalid integer value in CSV at (%d, %d) cell: %v", i, j, cell)
			}
		}
	}
	return nil
}
