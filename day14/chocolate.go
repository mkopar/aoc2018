package main

import (
	"fmt"
	"reflect"
	"strings"
	"aoc2018/lib/common"
)

func stringToIntArray(in string) []int {
	// a bit hackish, but for aoc is okay :)
	var tmpInt []int
	tmpStr := strings.Split(in, "")
	for _, el := range tmpStr {
		tmpInt = append(tmpInt, common.Atoi(el))
	}
	return tmpInt
}

func main() {
	input := 70283100 // input
	input2 := "702831"
	inputArray := stringToIntArray(input2)

	recipes := []int{3, 7} // first two recipes are 3 and 7
	pos1, pos2 := 0, 1     // position of each elf

	firstAppear := -1

	for i := 0; i < input+10; i++ {
		// check if inputArray already exists in recipes
		score := recipes[pos1] + recipes[pos2]
		if score >= 10 {
			recipes = append(recipes, score/10)
		}
		recipes = append(recipes, score%10)
		pos1, pos2 = (pos1+recipes[pos1]+1)%len(recipes), (pos2+recipes[pos2]+1)%len(recipes)
	}

	for i := 0; i < len(recipes)-len(inputArray); i++ {
		if reflect.DeepEqual(recipes[i:i+len(inputArray)], inputArray) {
			firstAppear = i
			break
		}
	}

	fmt.Println("Part 1:", recipes[input:input+10])
	fmt.Println("Part 2:", firstAppear)
}
