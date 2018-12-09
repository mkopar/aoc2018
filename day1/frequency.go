package main

import (
	"fmt"
	"aoc2018/lib/commonapi"
	"path/filepath"
	"log"
	"os"
	"bufio"
	"strconv"
)

func readToList(path string) []int {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tmp []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		} else {
			tmp = append(tmp, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tmp
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	changes := commonapi.ReadToList("day1/input")

	// puzzle1
	sum := 0
	for _, el := range changes {
		sum += el
	}
	fmt.Printf("Puzzle1 result: %d\n", sum)

	// puzzle2
	sum = 0
	pos := 0
	alreadySeenFreqs := []int{0}
	for {
		sum += changes[pos]
		if contains(alreadySeenFreqs, sum) {
			fmt.Printf("First frequency reached twice is: %d\n", sum)
			break
		} else {
			// check if change element already exists, if yes, change pos to the index, if no, pos++
			alreadySeenFreqs = append(alreadySeenFreqs, sum)
			pos = (pos + 1) % len(changes)
		}
	}
}
