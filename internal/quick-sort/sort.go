package quicksort

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

func QuickSort(input []int) []int {
	size := len(input)
	if size < 2 {
		return input
	}

	less := make([]int, 0, size)
	higher := make([]int, 0, size)
	middle := size / 2
	pivot := input[middle]

	for i := 0; i < size; i++ {
		if i == middle {
			continue
		}

		if input[i] <= pivot {
			less = append(less, input[i])
		} else {
			higher = append(higher, input[i])
		}
	}

	data := fmt.Sprintf("[QuickSort] [Report] number of elements: %v pivot %v | less %v | higher %v", size, pivot, less, higher)
	logger := logger.NewLogger()
	logger.Info("[QuickSort] [Input]", input)
	logger.Info(data, nil)

	numbers := append(
		append(QuickSort(less), pivot),
		QuickSort(higher)...,
	)

	logger.Info("[QuickSort] [Numbers]", numbers)

	return numbers
}
