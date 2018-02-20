package main

import (
	"context"
	"log"
	"os"
	"runtime"
)

func main() {
	maxConcurrency := runtime.NumCPU()
	runtime.GOMAXPROCS(maxConcurrency)

	args, err := parseArgs(os.Args)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	cmd, err := dispatchCmd(args)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	ctx := context.Background()
	cmd.Do(ctx)
}
