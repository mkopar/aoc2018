package main

import (
	"aoc2018/lib/common"
	"regexp"
	"fmt"
)

func parseInit(input string) string {
	r := regexp.MustCompile(`initial state: (.*)`)
	strParsed := r.FindStringSubmatch(input)[1:]
	return strParsed[0]
}

func parseMappings(input []string) (map[string]string) {
	tmp := make(map[string]string)
	r := regexp.MustCompile(`([\.#]{5}) => ([\.#])`)
	for _, l := range input {
		strParsed := r.FindStringSubmatch(l)[1:]
		oldval := strParsed[0]
		tmp[oldval] = strParsed[1]
	}
	return tmp
}

func getGenSum(in string, potzeroindex int) int {
	sum := 0
	for i := 0; i < len(in); i++ {
		if string(in[i]) == "#" {
			sum += i - potzeroindex // -3 * 20 generations
		}
	}
	return sum
}

func main() {
	strInput := common.ReadToStringList("day12/input")
	initState := parseInit(strInput[0])
	mappings := parseMappings(strInput[2:])

	generations := 200
	prevSum := getGenSum(initState, 0)
	prevDiff := 0
	for i := 0; i < generations; i++ {
		if i == 20 {
			fmt.Println("Part 1:", getGenSum(initState, 3*i))
		}
		initState = "..." + initState + "....."
		newState := ".."
		for j := 0; j < len(initState)-5; j++ {
			substr := initState[j : j+5]
			var replace string
			if _, ok := mappings[substr]; ok {
				// if substr exists in mappings, replace with mappings value for the key in next gen
				replace = mappings[substr]
			} else {
				// if substr does not exist in mappings, assume there is empty pot in next gen
				replace = "."
			}
			newState += replace
		}
		initState = newState
		curSum := getGenSum(initState, 3*i)
		curDiff := curSum - prevSum
		prevSum = curSum
		//fmt.Println(strconv.Itoa(i), curDiff)
		prevDiff = curDiff
	}
	part2calc := (50000000000 - generations) * prevDiff + getGenSum(initState, 3 * generations)
	fmt.Println("Part 2:", part2calc)
}
