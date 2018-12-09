package main

import (
	"time"
	"os"
	"log"
	"bufio"
	"strconv"
	"regexp"
	"reflect"
	"path/filepath"
	"sort"
	"fmt"
)

type guardSleeps struct {
	timestamp 		time.Time
	minutesAsleep	[]int
}

type shift struct {
	guardId			int
	shiftStart 		time.Time
	msg 			string
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

func parseInput(strShifts []string) []shift {
	var shifts []shift
	for _, strShft := range strShifts {
		r := regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})]\s+(.*)`)
		strParsed := r.FindStringSubmatch(strShft)[1:]
		timestamp, err := time.Parse("2006-01-02 15:04", strParsed[0])
		if err != nil {
			log.Fatal(err)
		}
		msgRe := regexp.MustCompile(`Guard #(\d+) begins shift|falls asleep|wakes up`)
		msg := msgRe.FindStringSubmatch(strParsed[1])
		if reflect.DeepEqual(msg, []string(nil)) {
			log.Fatal("invalid message, does not match")
		}
		guardId := 0
		if msg[1] != "" {
			guardId, err = strconv.Atoi(msg[1]) // msg[0] is all match (if any)
			if err != nil {
				log.Fatal(err)
			}
		}
		shifts = append(shifts, shift{guardId: guardId, shiftStart: timestamp, msg: msg[0]})
	}
	return shifts
}

func main() {
	path, err := filepath.Abs("day4/input")
	if err != nil {
		log.Fatal(err)
	}
	strShifts := readToList(path)
	shifts := parseInput(strShifts)
	sort.Slice(shifts, func(i, j int) bool {
		return shifts[i].shiftStart.Before(shifts[j].shiftStart)
	})

	// part 1
	guardId := shifts[0].guardId
	var sleepStart int
	var sleepEnd int
	var sleepMinutes []int
	guard := make(map[int][]int)
	mostSleepGuardId := 0
	mostSleep := 0
	for _, shift := range shifts {
		// shift.guardId = 0 is an event during shift
		if shift.guardId != 0 && guardId != shift.guardId {
			sleepStart = 0
			sleepEnd = 0
			guard[guardId] = append(guard[guardId], sleepMinutes...)
			if len(sleepMinutes) > mostSleep {
				mostSleep = len(sleepMinutes)
				mostSleepGuardId = guardId
			}
			sleepMinutes = []int(nil)
			// guard / shift change
			guardId = shift.guardId
			continue
		}
		// falls asleep
		if shift.msg == "falls asleep" {
			sleepStart = shift.shiftStart.Minute()
		}
		if shift.msg == "wakes up" {
			sleepEnd = shift.shiftStart.Minute()
			// generate sleep minutes
			for i := sleepStart; i < sleepEnd; i++ {
				sleepMinutes = append(sleepMinutes, i)
			}
		}
	}
	// need to add from last iteration
	guard[guardId] = append(guard[guardId], sleepMinutes...)

	// find most minute asleep
	mostMinuteAsleep := make(map[int]int)
	for _, el := range guard[mostSleepGuardId] {
		mostMinuteAsleep[el]++
	}

	minuteAsleepMax := 0
	minuteAsleepVal := 0
	for k, v := range mostMinuteAsleep {
		if minuteAsleepVal < v {
			minuteAsleepVal = v
			minuteAsleepMax = k
		}
	}

	fmt.Printf("Guard number %d spent most time sleeping\n", mostSleepGuardId)
	fmt.Printf("Most minute asleep %d\n", minuteAsleepMax)
	fmt.Printf("Part 1 result: %d\n", minuteAsleepMax * mostSleepGuardId)

	// part 2
	// find guard that is most frequently asleep on the same minute
	mostFrequentMinuteAsleep := map[int]map[int]int{}  // {"guardId": {"minute": "count"}}
	for grdId, minutes := range guard {
		mostFrequentMinuteAsleep[grdId] = map[int]int{}
		for _, min := range minutes {
			mostFrequentMinuteAsleep[grdId][min]++
		}
	}

	mostFrequentMinuteCount := 0
	mostFrequentMinute := 0
	mostFrequentMinuteGuard := 0
	for grdId, allMinCount := range mostFrequentMinuteAsleep {
		for minute, count := range allMinCount {
			if mostFrequentMinuteCount < count {
				mostFrequentMinuteCount = count
				mostFrequentMinuteGuard = grdId
				mostFrequentMinute = minute
			}
		}
	}
	fmt.Printf("Guard number %d slept most frequently on minute %d\n", mostFrequentMinuteGuard, mostFrequentMinute)
	fmt.Printf("Part 2 result: %d", mostFrequentMinuteGuard * mostFrequentMinute)
}
