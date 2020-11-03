package main

import "fmt"

// BubbleSort !
func BubbleSort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			n[i], n[i+1] = n[i+1], n[i]
		}
	}

	return n
}

// RecursiveBubbleSort !
func RecursiveBubbleSort(n []int, size int) []int {
	if size == 1 {
		return n
	}
	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			n[i], n[i+1] = n[i+1], n[i]
		}
	}

	RecursiveBubbleSort(n, size-1)

	return n
}

// InsertionSort !
func InsertionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := i; j >= 1 && n[j] < n[j-1]; j-- {
			n[j], n[j-1] = n[j-1], n[j]
		}
	}

	return n
}

func main() {
	n := []int{1, 39, 2, 9, 7, 54, 11}
	fmt.Println("BubbleSort:", BubbleSort(n))

	fmt.Println("RecursiveBubbleSort:", RecursiveBubbleSort(n, len(n)))
	fmt.Println("InsertionSort", InsertionSort(n))
}
