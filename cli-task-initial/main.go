package main

import (
	"os"

	"github.com/haquenafeem/practical-golang-tasks/runner"
)

func main() {
	args := os.Args
	runner.Init(args[1])
}
