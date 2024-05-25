package recursion

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

var log = logger.NewLogger()

func Factorial(num int) int {
	log.Info(fmt.Sprintf("Calculation %v", num), nil)

	if num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}
