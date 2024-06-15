package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Single element",
			input: []int{5},
			want:  []int{5},
		},
		{
			name:  "Two elements",
			input: []int{5, 4},
			want:  []int{4, 5},
		},
		{
			name:  "Three elements",
			input: []int{5, 1, 67},
			want:  []int{1, 5, 67},
		},
		{
			name:  "Four elements",
			input: []int{5, 1, 60, 0},
			want:  []int{0, 1, 5, 60},
		},
		{
			name:  "Five elements",
			input: []int{35, 55, 25, 95, 15},
			want:  []int{15, 25, 35, 55, 95},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
