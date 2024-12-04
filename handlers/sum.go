package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AddMatrixElems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Print each row
	sum := 0
	for _, row := range numericMatrix {
		for _, el := range row {
			sum += el
		}
	}
	fmt.Fprint(w, strconv.Itoa(sum)+"\n")
}
