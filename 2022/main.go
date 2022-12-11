package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//go:generate go run ./gen/

func main() {
	ctx := context.Background()
	err := run(ctx, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type runPart func(context.Context, []string) (string, error)

func run(ctx context.Context, args []string) error {
	currentDay := int(25 - time.Until(time.Date(2022, time.December, 25, 0, 0, 0, 0, time.Local)).Hours()/24)
	if len(args) >= 1 {
		currentDay, _ = strconv.Atoi(args[0])
	}
	part := "2"
	if len(args) >= 2 {
		part = args[1]
	}

	log.Printf("Runnning day %d part %s", currentDay, part)

	rp := runPartFactory(currentDay, part)
	answer, err := rp(ctx, args[2:])
	if err != nil {
		return fmt.Errorf("unable to run: %w", err)
	}

	fmt.Printf("\nAnswer: %s\n", answer)
	return nil
}
