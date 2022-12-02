package main

import (
	"context"
	"os"
	"log"
)

func main() {
	ctx := context.Background()
	err := run(ctx, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	return runDay2(ctx, args)
}

