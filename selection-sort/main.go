package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// SelectionSort performs selection sort on an array of integers
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// SimpleSort performs simple sort (bubble sort) on an array of integers
func SimpleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectionSortHandler(w http.ResponseWriter, r *http.Request) {
	arrStr := r.URL.Query().Get("arr")
	arr := parseIntArray(arrStr)
	if len(arr) == 0 {
		http.Error(w, "Invalid input array", http.StatusBadRequest)
		return
	}

	SelectionSort(arr)
	fmt.Fprintf(w, "Sorted array (Selection Sort): %v\n", arr)
}

func simpleSortHandler(w http.ResponseWriter, r *http.Request) {
	arrStr := r.URL.Query().Get("arr")
	arr := parseIntArray(arrStr)
	if len(arr) == 0 {
		http.Error(w, "Invalid input array", http.StatusBadRequest)
		return
	}

	SimpleSort(arr)
	fmt.Fprintf(w, "Sorted array (Simple Sort): %v\n", arr)
}

func parseIntArray(s string) []int {
	var arr []int
	strArr := strings.Split(s, ",")
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil
		}
		arr = append(arr, num)
	}
	return arr
}

func main() {
	http.HandleFunc("/selection-sort", selectionSortHandler)
	http.HandleFunc("/simple-sort", simpleSortHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
