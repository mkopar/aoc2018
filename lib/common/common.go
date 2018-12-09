package common

import (
	"os"
	"log"
	"bufio"
	"strconv"
)

func ReadToStringList(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	var tmp []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Check(err)
		tmp = append(tmp, scanner.Text())
	}

	err = scanner.Err()
	Check(err)

	return tmp
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Atoi(tmp string) int {
	out, err := strconv.Atoi(tmp)
	Check(err)
	return out
}

func IntListContains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ParseStringListToIntList(input []string) []int {
	var tmp []int
	for _, el := range input {
		intEl := Atoi(el)
		tmp = append(tmp, intEl)
	}
	return tmp
}
