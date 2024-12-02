package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func EchoMatrix(w http.ResponseWriter, r *http.Request) {
	matrix, err := getMatrixFromFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	// Print each row
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

// Parses the file from the request into a matrix
func getMatrixFromFile(r *http.Request) ([][]string, error) {
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
