package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestEchoMatrix(t *testing.T) {
	testCases := []struct {
		name               string
		method             string
		csvContent         string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:   "Valid CSV square matrix",
			method: http.MethodPost,
			csvContent: `1,2,3
4,5,6
7,8,9`,
			expectedStatusCode: http.StatusOK,
			expectedBody:       "1,2,3\n4,5,6\n7,8,9\n",
		},
		{
			name:   "Invalid CSV rectangle matrix",
			method: http.MethodPost,
			csvContent: `1,2,3
4,5,6
7,8,9
10,11,12`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "number of rows not equal to size of row",
		},
		{
			name:               "Invalid method",
			method:             http.MethodGet,
			csvContent:         "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "request Content-Type isn't multipart/form-data",
		},
		{
			name:   "Invalid CSV with non-numeric value",
			method: http.MethodPost,
			csvContent: `21,2,685
			47594763,five,859
			75,18,9285`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "invalid integer value in CSV",
		},
		{
			name:               "Empty CSV",
			method:             http.MethodPost,
			csvContent:         "",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "file is empty",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var req *http.Request
			if testCase.method == http.MethodPost {
				req, _ = createMultipartRequest(testCase.csvContent)
			} else {
				req = httptest.NewRequest(testCase.method, "/echo", nil)
			}

			rr := httptest.NewRecorder()
			EchoMatrix(rr, req)

			// Check status code
			if rr.Code != testCase.expectedStatusCode {
				t.Errorf("Unexpected status code: got %v, want %v", rr.Code, testCase.expectedStatusCode)
			}

			// Check response body
			if !strings.Contains(rr.Body.String(), testCase.expectedBody) {
				t.Errorf("Unexpected response body: got %v, want substring %v", rr.Body.String(), testCase.expectedBody)
			}
		})
	}
}
