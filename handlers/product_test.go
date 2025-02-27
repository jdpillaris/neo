package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMultiplyMatrixElems(t *testing.T) {
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
			expectedBody:       "362880\n",
		},
		{
			name:   "Valid matrix with zero",
			method: http.MethodPost,
			csvContent: `1,2,3,4
5,6,7,8
9,10,0,11
12,13,14,15`,
			expectedStatusCode: http.StatusOK,
			expectedBody:       "0\n",
		},
		{
			name:   "Invalid CSV with extra cell",
			method: http.MethodPost,
			csvContent: `1,2,38
			4,5,6,7
			7,8,926`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "wrong number of fields",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var req *http.Request
			if testCase.method == http.MethodPost {
				req, _ = createMultipartRequest(testCase.csvContent)
			} else {
				req = httptest.NewRequest(testCase.method, "/multiply", nil)
			}

			rr := httptest.NewRecorder()
			MultiplyMatrixElems(rr, req)

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
