package binarysearch_test

import (
	"testing"

	binarysearch "github.com/silasstoffel/grokking-algorithms/internal/binary-search"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	t.Run("Should find item int the middle of the list", func(t *testing.T) {
		search := 9
		numbers := []int{1, 2, 3, 4, 5, 7, 9, 11}
		output := binarysearch.Search(numbers, search)

		if output.Index < 0 {
			t.Errorf("Expect index more or equal than 0 and receive %v", output.Index)
		}

		if output.IterationToSearch != 3 {
			t.Errorf("Expect iteration to search %v, got %v", 3, output.IterationToSearch)
		}

		if numbers[output.Index] != search {
			t.Errorf("Expect number %v got %v", search, numbers[output.Index])
		}
	})

	t.Run("Should find the first item", func(t *testing.T) {
		search := 1
		numbers := []int{1, 2, 3, 4, 5, 7, 9, 11}
		output := binarysearch.Search(numbers, search)

		if output.Index < 0 {
			t.Errorf("Expect index more or equal than 0 and receive %v", output.Index)
		}

		if output.IterationToSearch != 0 {
			t.Errorf("Expect iteration to search %v, got %v", 0, output.IterationToSearch)
		}

		if numbers[output.Index] != search {
			t.Errorf("Expect number %v got %v", search, numbers[output.Index])
		}
	})

	t.Run("Should find the last item", func(t *testing.T) {
		search := 11
		numbers := []int{1, 2, 3, 4, 5, 7, 9, 11}
		output := binarysearch.Search(numbers, search)

		if output.Index < 0 {
			t.Errorf("Expect index more or equal than 0 and receive %v", output.Index)
		}

		if output.IterationToSearch != 0 {
			t.Errorf("Expect iteration to search %v, got %v", 0, output.IterationToSearch)
		}

		if numbers[output.Index] != search {
			t.Errorf("Expect number %v got %v", search, numbers[output.Index])
		}
	})

	t.Run("Should return empty when item does not exists in the list", func(t *testing.T) {
		search := 25
		numbers := []int{1, 2, 3, 4, 5, 7, 9, 11}
		output := binarysearch.Search(numbers, search)

		if output.Index != -1 {
			t.Errorf("Expect index equal than %v and receive %v", -1, output.Index)
		}

		if output.IterationToSearch == 0 {
			t.Errorf("Expect iteration to search more than 0")
		}
	})
}
