package main

import (
	"fmt"
	"strings"
	"aoc2018/lib/common"
)

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
	polymer := common.ReadToStringList("day5/input")[0] // just one element
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
