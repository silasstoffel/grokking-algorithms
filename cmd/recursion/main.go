package main

import (
	"github.com/silasstoffel/grokking-algorithms/internal/logger"
	"github.com/silasstoffel/grokking-algorithms/internal/recursion"
)

var log = logger.NewLogger()

func main() {
	f := recursion.Factorial(5)
	log.Info("Factorial: ", f)
}
