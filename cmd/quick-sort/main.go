package main

import (
	"github.com/silasstoffel/grokking-algorithms/internal/logger"
	quicksort "github.com/silasstoffel/grokking-algorithms/internal/quick-sort"
)

func main() {
	logger := logger.NewLogger()
	numbers := []int{98, 88, 8, 78, 28, 38, 68, 58}
	logger.Info("Numbers to sort", numbers)
	sorted := quicksort.QuickSort(numbers)
	logger.Info("Sorted numbers", sorted)
}
