package selectionsort

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

func MinIndex(items []int, start int) int {
	minIndex := start
	for i := start + 1; i < len(items); i++ {
		if items[i] < items[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func Sort(items []int) []int {
	logger := logger.NewLogger()
	counter := len(items)
	sorted := make([]int, counter)
	copy(sorted, items)

	for i := 0; i < counter-1; i++ {
		minIndex := MinIndex(sorted, i)
		logger.Info(fmt.Sprintf("Min value %v - Min index: %v", items[i], minIndex), nil)
		sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
	}

	return sorted
}
