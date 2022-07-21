//go:build unit
// +build unit

package math1_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"testingpractice/math1"
)

func TestDivide(t *testing.T) {
	result := math1.Divide(2, 2)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestAdd(t *testing.T) {
	sum := math1.Add(1, 2)
	if condition := sum != 3; condition {
		t.Errorf("Expected 3, got %d", sum)
	} else {
		t.Logf("Expected 3, got %d", sum)
		t.Log("Test passed")
	}
}

func TestTableAdd(t *testing.T) {
	var tests = []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
	}
	for _, test := range tests {
		if got := math1.Add(test.a, test.b); got != test.c {
			t.Errorf("Add(%d, %d) = %d, want %d", test.a, test.b, got, test.c)
		}
	}
	t.Log("Test passed")
}

func TestReadFile(t *testing.T) {
	file := "test.txt"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Error reading file %s", err)
	}
	if string(data) != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", string(data))
	}
	t.Logf("File %s read successfully", file)
}

func TestHttpRequest(t *testing.T) {
	hadler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Hello World")
	}
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	w := httptest.NewRecorder()
	hadler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 400 != resp.StatusCode {
		t.Errorf("Expected status code 400, got %d", resp.StatusCode)
	}
}
