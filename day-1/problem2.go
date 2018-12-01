package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	// Read values
	list, err := readList(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Map of frequencies already seen
	seen := make(map[int64]struct{})
	var done bool
	var f int64

	for !done {
		for _, v := range list {
			f += v
			if _, found := seen[f]; found {
				done = true
				break
			}
			seen[f] = struct{}{}
		}
	}

	log.Print(f)
}

func readList(r io.Reader) ([]int64, error) {
	var list []int64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
