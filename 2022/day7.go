package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func runDay7Part1(ctx context.Context, args []string) (string, error) {
	path := "day7.input"
	if len(args) > 0 {
		path = args[0]
	}
	rootDir, err := readInputDay7(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	dirTotals := map[string]int{}
	rootDir.Walk(func(path []string, dir directory) {
		dirTotals[strings.Join(path, "/")] = dir.TotalSize()
	})

	candidateTotal := 0
	for _, dir := range dirTotals {
		if dir > 100000 {
			continue
		}
		candidateTotal += dir
	}

	return fmt.Sprintf("%d", candidateTotal), nil
}

func runDay7Part2(ctx context.Context, args []string) (string, error) {
	path := "day7.input"
	if len(args) > 0 {
		path = args[0]
	}
	_, err := readInputDay7(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	panic("not implemented")
}

type directory struct {
	Name        string
	Files       map[string]int
	Directories map[string]*directory
}

func newDirectory(name string) *directory {
	return &directory{
		Name:        name,
		Files:       map[string]int{},
		Directories: map[string]*directory{},
	}
}

func (d *directory) TotalSize() int {
	total := 0
	for _, size := range d.Files {
		total += size
	}
	for _, dir := range d.Directories {
		total += dir.TotalSize()
	}
	return total
}

func (d directory) Walk(walker func([]string, directory)) {
	path := []string{d.Name}
	walker(path, d)
	for _, dir := range d.Directories {
		dir.Walk(func(childPath []string, childDir directory) {
			walker(append(path, childPath...), childDir)
		})
	}
}

func readInputDay7(inputFile string) (*directory, error) {
	input, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open input: %w", err)
	}
	defer input.Close()

	var rootDir *directory
	var path []*directory

	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		panic("empty input file")
	}
	// post check loop due to inner loop for ls
CMDScanLoop:
	for {
		line := scanner.Text()

		if !strings.HasPrefix(line, "$") {
			panic(fmt.Sprintf("unexpected command: %q", line))
		}

		switch cmd := line[2:]; {
		default:
			panic(fmt.Sprintf("unknown command: %q", cmd))
		case cmd == "cd /":
			if rootDir == nil {
				rootDir = newDirectory("/")
			}
			path = []*directory{rootDir}
		case cmd == "cd ..":
			_, path = pop(path)
		case strings.HasPrefix(cmd, "cd "):
			cwd := path[len(path)-1]
			to := cmd[3:]

			// is this even possible? should it already have been encounted via ls?
			if _, ok := cwd.Directories[to]; !ok {
				newDir := newDirectory(to)
				cwd.Directories[to] = newDir
			}

			toDir := cwd.Directories[to]
			path = push(path, toDir)
		case cmd == "ls":
			cwd := path[len(path)-1]
		LSScanLoop:
			for scanner.Scan() {
				line = scanner.Text()

				// this is a new command, go back to outer loop
				if strings.HasPrefix(line, "$") {
					continue CMDScanLoop
				}

				if strings.HasPrefix(line, "dir ") {
					dirName := line[4:]
					cwd.Directories[dirName] = newDirectory(dirName)
					continue LSScanLoop
				}

				var name string
				var size int
				_, err := fmt.Sscanf(line, "%d %s", &size, &name)
				if err != nil {
					return nil, fmt.Errorf("unable to parse line %q: %w", line, err)
				}

				cwd.Files[name] = size
			}
		}

		if !scanner.Scan() {
			break CMDScanLoop
		}
	}

	return rootDir, nil
}
