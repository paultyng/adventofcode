package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

//go:generate go run ./gen/

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ctx := context.Background()
	err := run(ctx, flag.Args())
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

	fmt.Printf("Runnning day %d part %s\n", currentDay, part)

	rp := runPartFactory(currentDay, part)
	answer, err := rp(ctx, args[2:])
	if err != nil {
		return fmt.Errorf("unable to run: %w", err)
	}

	fmt.Printf("\nAnswer: %s\n", answer)
	return nil
}
