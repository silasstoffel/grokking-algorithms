package recursion_test

import (
	"testing"

	"github.com/silasstoffel/grokking-algorithms/internal/recursion"
)

func TestFactorial(t *testing.T) {
	t.Parallel()

	t.Run("Should return 1", func(t *testing.T) {
		f := recursion.Factorial(1)
		if f != 1 {
			t.Errorf("Expected 1 and got %v", f)
		}
	})

	t.Run("Should calculate the factorial number", func(t *testing.T) {
		f := recursion.Factorial(5)
		if f != 120 {
			t.Errorf("Expected 120 and got %v", f)
		}
	})
}
