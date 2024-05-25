package recursion

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

var log = logger.NewLogger()

func Factorial(num int) int {
	if num == 1 {
		log.Info(fmt.Sprintf("Calculation !%v = %v", 1, num), nil)
		return 1
	}
	r := num * Factorial(num-1)
	log.Info(fmt.Sprintf("Calculation !%v = %v", num, r), nil)
	return r
}
