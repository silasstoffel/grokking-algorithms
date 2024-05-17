package binarysearch

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

type BinarySearchOutput struct {
	Index             int
	IterationToSearch int
}

func Search(list []int, search int) *BinarySearchOutput {
	output := &BinarySearchOutput{
		Index:             -1,
		IterationToSearch: 0,
	}
	counter := len(list)
	logger := logger.NewLogger()

	if counter == 0 {
		logger.Info("List argument contains 0 items", nil)
		return output
	}

	if list[0] == search {
		logger.Info("Search value is on first element", nil)
		output.Index = 0
		return output
	}

	// last index
	right := counter - 1

	if list[right] == search {
		output.Index = right
		logger.Info("Search value is on last element", nil)
		return output
	}

	logger.Info("Starting iteration to search item", nil)
	left := 0
	logger.Info(fmt.Sprintf("Searching: %v", search), nil)

	for left <= right {
		output.IterationToSearch++
		middle := (left + right) / 2
		currentValue := list[middle]

		logger.Info(fmt.Sprintf("Current Index: %v - Current Value: %v - Iteration: %v", middle, currentValue, output.IterationToSearch), nil)

		if currentValue == search {
			output.Index = middle
			logger.Info("Found item", nil)
			return output
		}

		if currentValue > search {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return output
}
