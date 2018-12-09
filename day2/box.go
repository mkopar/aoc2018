package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"path/filepath"
)

func readToList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tmp []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		} else {
			tmp = append(tmp, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tmp
}

func main() {
	path, err := filepath.Abs("day2/input")
	if err != nil {
		log.Fatal(err)
	}
	boxIDs := readToList(path)

	// part1
	m := make(map[string]int)
	for _, element := range boxIDs {
		twosamevalid := true
		threesamevalid := true
		for _, chr := range element {
			tmpcount := strings.Count(element, string(chr))
			if tmpcount == 2 && twosamevalid {
				m["two"]++
				twosamevalid = false
			}
			if tmpcount == 3 && threesamevalid {
				m["three"]++
				threesamevalid = false
			}
			if !twosamevalid && !threesamevalid {
				break
			}
		}
	}
	fmt.Println(m["two"] * m["three"])

	// part2
	for i := 0; i < len(boxIDs); i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			differpos := 0
			changes := 0
			for idx := range boxIDs[j] {
				if boxIDs[i][idx] != boxIDs[j][idx] {
					changes++
					differpos = idx
				}
				if changes > 1 {
					break
				}
			}
			if changes == 1 {
				fmt.Printf("Differing in one character at position %d\n", differpos)
				fmt.Println(boxIDs[i])
				fmt.Println(boxIDs[j])
				fmt.Printf("Result: %s%s", boxIDs[i][:differpos], boxIDs[i][differpos+1:])
				fmt.Println()
			}
		}
	}
}
