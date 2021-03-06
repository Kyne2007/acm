package main

import (
	"fmt"
	"net/http"
)

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	n := len(matrix)
	m := len(matrix[0])

	for i, j := 0, m-1; i < n && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}

type ha struct {
}

type haer interface {
	wuha()
}

func (h *ha) wuha() {
	fmt.Println("wuha!")
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello wuha\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4444", nil)

}
