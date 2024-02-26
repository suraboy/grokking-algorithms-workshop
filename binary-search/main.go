package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// GenerateSortedArray generates a sorted array of integers from 1 to n
func GenerateSortedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

// BinarySearch performs binary search on a sorted array
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid // found
		} else if arr[mid] < target {
			low = mid + 1 // search in the right half
		} else {
			high = mid - 1 // search in the left half
		}
	}
	return -1 // not found
}

// SimpleSearch performs simple linear search
func SimpleSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // found
		}
	}
	return -1 // not found
}

func main() {
	// Generate a sorted array with 4 million elements
	arr := GenerateSortedArray(400000000)

	// Define HTTP handler functions
	http.HandleFunc("/binary-search", func(w http.ResponseWriter, r *http.Request) {
		target, err := strconv.Atoi(r.URL.Query().Get("target"))
		if err != nil {
			http.Error(w, "Invalid target", http.StatusBadRequest)
			return
		}
		index := BinarySearch(arr, target)
		fmt.Fprintf(w, "Binary search result: %d\n", index)
	})

	http.HandleFunc("/simple-search", func(w http.ResponseWriter, r *http.Request) {
		target, err := strconv.Atoi(r.URL.Query().Get("target"))
		if err != nil {
			http.Error(w, "Invalid target", http.StatusBadRequest)
			return
		}
		index := SimpleSearch(arr, target)
		fmt.Fprintf(w, "Simple search result: %d\n", index)
	})

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
