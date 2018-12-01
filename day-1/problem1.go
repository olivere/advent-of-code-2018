package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)

	var f int64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		f += v
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print(f)
}
