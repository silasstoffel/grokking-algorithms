package main

import (
	"github.com/silasstoffel/grokking-algorithms/internal/logger"
	selectionsort "github.com/silasstoffel/grokking-algorithms/internal/selection-sort"
)

func main() {
	logger := logger.NewLogger()

	numbers := []int{20, 10, 15, 7, 3, 28, 1}

	logger.Info("Sorting number", numbers)

	sorted := selectionsort.Sort(numbers)

	logger.Info("Sorted numbers", sorted)
}
