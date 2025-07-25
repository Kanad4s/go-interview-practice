package main

import (
	"fmt"
	"slices"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)

	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, val := range numbers {
		if val > max {
			max = val
		}
	}
	return max
	// check for empty slice
	// init with first value
	// check subsequent vals for greater value
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
	res := make([]int, 0, len(numbers))

	for _, v := range numbers {
		if !slices.Contains(res, v) {
			res = append(res, v)
		}
	}

	return res
	// allocate result slice
	// loop thru slice filtering out dups
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
	res := make([]int, len(slice))
	for idx, val := range slice {
		res[len(res)-1-idx] = val
	}
	return res
	// allocate result slice
	// loop thru slice filling res from the end backwards
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	res := make([]int, 0, len(numbers))
	for _, val := range numbers {
		if val%2 == 0 {
			res = append(res, val)
		}
	}
	return res
	// allocate result slice
	// loop thru slice filtering out dups
}
