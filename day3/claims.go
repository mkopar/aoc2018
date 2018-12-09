package main

import (
	"os"
	"log"
	"bufio"
	"path/filepath"
	"regexp"
	"strconv"
	"fmt"
)

type claim struct {
	id		int
	left	int
	top		int
	width	int
	height	int
}

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

func parseStringToInt(input []string) []int {
	var output []int
	for _, el := range input {
		val, err := strconv.Atoi(el)
		if err != nil {
			log.Fatal(err)
		} else {
			output = append(output, val)
		}
	}
	return output
}

func parseInput(strClaims []string) []claim {
	var claims []claim
	for _, strClm := range strClaims {
		r := regexp.MustCompile(`#(\d+)\s+@\s+(\d+),(\d+):\s+(\d+)x(\d+)`)
		strParsed := r.FindStringSubmatch(strClm)[1:]
		strParsedToInt := parseStringToInt(strParsed)
		claims = append(claims, claim{strParsedToInt[0], strParsedToInt[1], strParsedToInt[2], strParsedToInt[3], strParsedToInt[4]})
	}
	return claims
}

func main() {
	path, err := filepath.Abs("day3/input")
	if err != nil {
		log.Fatal(err)
	}
	strClaims := readToList(path)
	claims := parseInput(strClaims)

	// part 1
	const squareLen = 1000
	var square [squareLen][squareLen]int
	var overlap [1233]bool // on the appropriate index there will be overlap - checked input file and know that there are 1233 ids - should also automate it by getting max id but I'm lazy :)

	for _, clm := range claims {
		for i := 0; i < clm.height; i++ {
			for j := 0; j < clm.width; j++ {
				if square[clm.top + i][clm.left + j] == 0 {
					square[clm.top + i][clm.left + j] = clm.id
				} else {
					overlap[clm.id - 1] = true
					if square[clm.top + i][clm.left + j] - 1 >= 0 {
						overlap[square[clm.top + i][clm.left + j] - 1] = true
					}
					square[clm.top + i][clm.left + j] = -1
				}
			}
		}
	}
	duplicates := 0
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square[0]); j++ {
			if square[i][j] == -1 {
				duplicates++
			}
		}
	}
	fmt.Printf("Part 1: duplicates in %d square inches\n", duplicates)

	for idx, ovrlp := range overlap {
		if !ovrlp {
			fmt.Printf("Part 2: only not overlapping id is %d\n", idx+1)
		}
	}
}
