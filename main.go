package main

import (
	"fmt"
	"os"

	generator "github.com/robertsong9972/gorm-gen/internal"
)

const (
	Gen = "gen"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic!", err)
		}
	}()
	tool := os.Args[1]
	switch tool {
	case Gen:
		generator.Gen()
	default:
		fmt.Println("unknown tool type, supported tool types are:")
		fmt.Println(Gen)
	}
}
