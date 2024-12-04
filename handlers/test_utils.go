package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
)

func createMultipartRequest(csvContent string) (*http.Request, string) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, strings.NewReader(csvContent))
	if err != nil {
		panic(err)
	}

	writer.Close()
	req := httptest.NewRequest(http.MethodPost, "/upload", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, writer.FormDataContentType()
}
