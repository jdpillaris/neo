package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

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

func getIntMatrix(grid [][]string) ([][]int, error) {
	gridLength := len(grid)
	if gridLength == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	var numericMatrix [][]int
	for i, row := range grid {
		if len(row) != gridLength {
			return nil, fmt.Errorf("number of rows not equal to size of row %d", i)
		}
		var numericRow []int
		for j, cell := range row {
			val, err := strconv.Atoi(cell)
			if err != nil {
				return nil, fmt.Errorf("invalid integer value in CSV at (%d, %d) cell: %v", i, j, cell)
			}
			numericRow = append(numericRow, val)
		}
		numericMatrix = append(numericMatrix, numericRow)
	}
	return numericMatrix, nil
}
