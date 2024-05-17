package main

import (
	binarysearch "github.com/silasstoffel/grokking-algorithms/internal/binary-search"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 7, 9, 11, 12, 14, 16, 18, 20, 21, 23, 25}

	search := 21
	binarysearch.Search(numbers, search)

	search = 3
	binarysearch.Search(numbers, search)

	search = 11
	binarysearch.Search(numbers, search)

	search = 1
	binarysearch.Search(numbers, search)

	search = 25
	binarysearch.Search(numbers, search)
}
