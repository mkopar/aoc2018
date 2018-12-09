package main

import (
	"fmt"
	"aoc2018/lib/common"
)

func main() {
	changesStr := common.ReadToStringList("day1/input")
	changes := common.ParseStringListToIntList(changesStr)

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
		if common.IntListContains(alreadySeenFreqs, sum) {
			fmt.Printf("Part 2 result: %d\n", sum)
			break
		} else {
			// check if change element already exists, if yes, change pos to the index, if no, pos++
			alreadySeenFreqs = append(alreadySeenFreqs, sum)
			pos = (pos + 1) % len(changes)
		}
	}
}
