package main

import (
	"aoc2018/lib/common"
	"fmt"
	"regexp"
	"container/list"
)

func parse(in string) (int, int) {
	r := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`)
	strParsed := r.FindStringSubmatch(in)[1:]
	players := common.Atoi(strParsed[0])
	lastMarbleWorth := common.Atoi(strParsed[1])

	return players, lastMarbleWorth
}

func max(a []int) int {
	high := 0
	for _, el := range a {
		if el > high {
			high = el
		}
	}
	return high
}

func part1(players int, lastMarbleWorth int) int {
	playerScore := make([]int, players)
	playingGrid := list.New()
	cur := playingGrid.PushFront(0)
	for marbleWorth := 1; marbleWorth < lastMarbleWorth; marbleWorth++ {
		if marbleWorth%23 == 0 {
			// go back for 6 elements
			for i := 0; i < 6; i++ {
				if cur = cur.Prev(); cur == nil {
					cur = playingGrid.Back()
				}
			}
			playerScore[marbleWorth%players] += marbleWorth + cur.Prev().Value.(int)
			playingGrid.Remove(cur.Prev())
		} else {
			if cur = cur.Next(); cur == nil {
				cur = playingGrid.Front()
			}
			cur = playingGrid.InsertAfter(marbleWorth, cur)
		}
	}
	highScore := max(playerScore)
	return highScore
}

func main() {
	strInput := common.ReadToStringList("day9/input")[0]
	players, last := parse(strInput)
	result1 := part1(players, last)
	fmt.Println("Part 1 result:", result1)
	result2 := part1(players, last*100)
	fmt.Println("Part 2 result:", result2)
}
