package main

import (
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/bfs"
	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

/**
 * Breadth First Search
 *
 * Mango Seller case.
 * The objective is find the mango seller.
 *
 * Schema used is
 * person -> friends
 *
 * me     -> [bob, claire, alice]
 * bob    -> [anju, peggy]
 * claire -> [thom, jonny]
 * alice  -> [peggy]
 */

var sellers = map[string][]string{
	"me":     {"bob", "claire", "alice"},
	"bob":    {"anju", "peggy"},
	"claire": {"thom", "jonny"},
	"alice":  {"peggy"},
	"anju":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

var log = logger.NewLogger()

func main() {
	log.Info("[main] starting search to get mongo seller", nil)
	firstSeller := "me"
	seller, err := bfs.Bfs(firstSeller, sellers)

	if err != nil {
		log.Info(fmt.Sprintf("[main] Not found mongo seller starting by %v", firstSeller), nil)
	}

	log.Info(fmt.Sprintf("[main] Mango seller found. His name is %v", seller), nil)

	log.Info("[main] finishing search to got mongo seller", nil)
}
