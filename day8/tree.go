package main

import (
	"aoc2018/lib/common"
	"strings"
	"fmt"
)

func part1(tree []int, curIdx int) (int, int) {
	sum := 0
	childs, metadata := tree[curIdx], tree[curIdx+1]
	for {
		if childs == 0 {
			for i := curIdx + 2; i < metadata+curIdx+2; i++ {
				sum += tree[i]
			}
			break
		} else {
			tmp, tmpIdx := part1(tree, curIdx+2)
			curIdx = tmpIdx
			sum += tmp
			childs--
		}
	}
	return sum, curIdx + metadata
}

func part2(tree []int, curIdx int) (int, int) {
	sum := 0
	childs, metadata := tree[curIdx], tree[curIdx+1]
	var childVal []int
	childCount := 0
	for {
		if childs == 0 {
			for i := curIdx + 2; i < metadata+curIdx+2; i++ {
				sum += tree[i]
			}
		} else {
			tmp, tmpIdx := part2(tree, curIdx+2)
			childVal = append(childVal, tmp)
			curIdx = tmpIdx
		}
		childCount++
		if childCount >= childs {
			break
		}
	}
	if len(childVal) > 0 {
		for i := curIdx + 2; i < metadata+curIdx+2; i++ {
			if tree[i]-1 < len(childVal) {
				sum += childVal[tree[i]-1]
			}
		}
	}
	return sum, curIdx + metadata
}

func main() {
	strInput := common.ReadToStringList("day8/input")[0]
	parsed := common.ParseStringListToIntList(strings.Fields(strInput))
	result1, _ := part1(parsed, 0)
	fmt.Println("Part 1 result:", result1)
	result2, _ := part2(parsed, 0)
	fmt.Println("Part 2 result:", result2)
}
