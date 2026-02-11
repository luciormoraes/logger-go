package main

import (
	"fmt"

	"github.com/luciormoraes/logger-go/pocketlog"
)

func main() {
	lvl := pocketlog.LevelDebug
	fmt.Printf("Level: %v\n", lvl)
}
