package bfs

import (
	"errors"
	"fmt"

	"github.com/silasstoffel/grokking-algorithms/internal/logger"
)

var log = logger.NewLogger()

func isMangoSeller(seller string) bool {
	lastLetter := seller[len(seller)-1]
	return string(lastLetter) == "e"
}

func Bfs(person string, sellers map[string][]string) (string, error) {
	evaluated := make(map[string]bool)
	queue := make([]string, len(sellers[person]))
	queue = append(queue, sellers[person]...)

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if name == "" {
			continue
		}

		log.Info(fmt.Sprintf("Verifying %v", name), nil)

		if _, ok := evaluated[name]; ok {
			log.Info("Name already verified.", nil)
			continue
		}

		evaluated[name] = true

		if isMangoSeller(name) {
			log.Info(fmt.Sprintf("Found mango seller. Mango seller name is %v", name), nil)
			return name, nil
		}
		if len(sellers[name]) > 0 {
			queue = append(queue, sellers[name]...)
		}
		log.Info(fmt.Sprintf("Mango seller(%v) not found", name), nil)
	}

	return "", errors.New("seller not found")
}
