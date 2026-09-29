package main

import (
	"fmt"
	"os"

	"github.com/Spider-has/task-1/internal/runner"
)

func main() {
	if err := runner.Run(os.Stdin, os.Stdout); err != nil {
		fmt.Println(err)
	}
}
