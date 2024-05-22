package selectionsort_test

import (
	"testing"

	selectionsort "github.com/silasstoffel/grokking-algorithms/internal/selection-sort"
)

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	t.Run("Should sort the list", func(t *testing.T) {
		numbers := []int{11, 7, 5, 9}
		sorted := selectionsort.Sort(numbers)

		if sorted[0] != 5 {
			t.Errorf("Expect %v and received %v at 0 index", 5, sorted[0])
		}

		if sorted[3] != 11 {
			t.Errorf("Expect %v and received %v at 3 index", 11, sorted[3])
		}
	})

}
