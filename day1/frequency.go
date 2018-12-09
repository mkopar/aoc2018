package main

import (
	"fmt"
	"aoc2018/lib/common"
	"strconv"
	"log"
)

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func parseStringToInt(input []string) []int {
	var tmp []int
	for _, el := range input {
		intEl, err := strconv.Atoi(el)
		if err != nil {
			log.Fatal(err)
		}
		tmp = append(tmp, intEl)
	}
	return tmp
}

func main() {
	changesStr := common.ReadToStringList("day1/input")
	changes := parseStringToInt(changesStr)

	// puzzle1
	sum := 0
	for _, el := range changes {
		sum += el
	}
	fmt.Printf("Part 1 result: %d\n", sum)

	// puzzle2
	sum = 0
	pos := 0
	alreadySeenFreqs := []int{0}
	for {
		sum += changes[pos]
		if contains(alreadySeenFreqs, sum) {
			fmt.Printf("Part 2 result: %d\n", sum)
			break
		} else {
			// check if change element already exists, if yes, change pos to the index, if no, pos++
			alreadySeenFreqs = append(alreadySeenFreqs, sum)
			pos = (pos + 1) % len(changes)
		}
	}
}
