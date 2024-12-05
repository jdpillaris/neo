package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func MultiplyMatrixElems(w http.ResponseWriter, r *http.Request) {
	grid, err := getGridFromFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var numericMatrix [][]int
	numericMatrix, err = getIntMatrix(grid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Multiply integers
	product := 1
	for _, row := range numericMatrix {
		for _, el := range row {
			if el == 0 {
				fmt.Fprint(w, "0\n")
				return
			}
			product *= el
		}
	}
	fmt.Fprint(w, strconv.Itoa(product)+"\n")
}
