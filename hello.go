package main

import (
	"fmt"
)

func findNumberWithoutNegativeEquivalent(numbers []int) int {
	result := 0

	for _, num := range numbers {
		// Проверяем, есть ли отрицательное аналогичное число в массиве
		if !contains(numbers, -num) {
			result = num
			break
		}
	}

	return result
}

func contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{-1, 1, -2, 2, 3, -3, 4, 0}

	result := findNumberWithoutNegativeEquivalent(numbers)

	fmt.Printf("Число без отрицательного аналога: %d\n", result)
}
