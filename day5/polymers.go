package main

import (
	"os"
	"log"
	"bufio"
	"path/filepath"
	"fmt"
	"strings"
)

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tmp string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		} else {
			tmp = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tmp
}

func reactPolymer(polymer string) string {
	polymerCheck := polymer
	for {
		for i := 0; i < len(polymer) - 1; i++ {
			if (polymer[i] == polymer[i+1] + 32) || (polymer[i] + 32 == polymer[i+1]) {
				polymer = polymer[:i] + polymer[i+2:]
				break
			}
		}
		if len(polymer) == len(polymerCheck) {
			break
		} else {
			polymerCheck = polymer
		}
	}
	return polymer
}

func main() {
	path, err := filepath.Abs("day5/input")
	if err != nil {
		log.Fatal(err)
	}
	polymer := readFile(path)
	// part one
	reactedPolymer := reactPolymer(polymer)
	fmt.Println("Part 1 result:", len(reactedPolymer))

	minLen := len(reactedPolymer)
	// part two - using capitals
	for i := 65; i < 90; i++ {
		tmpPolymer := strings.Replace(polymer, string(i), "",-1)
		tmpPolymer = strings.Replace(tmpPolymer, string(i+32), "",-1)
		tmpReactedPolymer := reactPolymer(tmpPolymer)
		if len(tmpReactedPolymer) < minLen {
			minLen = len(tmpReactedPolymer)
		}
	}
	fmt.Println("Part 2 result:", minLen)
}
