package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	err := run(ctx, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	answer, err := runDay2Part2(ctx, args)
	if err != nil {
		return fmt.Errorf("unable to run: %w", err)
	}

	fmt.Printf("\nAnswer: %s\n", answer)
	return nil
}
