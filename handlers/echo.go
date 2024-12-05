package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func EchoMatrix(w http.ResponseWriter, r *http.Request) {
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

	// Print each row
	var response string
	for _, row := range grid {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
