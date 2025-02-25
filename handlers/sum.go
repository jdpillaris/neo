package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AddMatrixElems(w http.ResponseWriter, r *http.Request) {
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

	// Add all integers
	sum := 0
	for _, row := range numericMatrix {
		for _, el := range row {
			sum += el
		}
	}
	fmt.Fprint(w, strconv.Itoa(sum)+"\n")
}
