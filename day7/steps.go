package main

import (
	"aoc2018/lib/common"
	"regexp"
	"fmt"
	"sort"
)

type stp struct {
	name   string
	before string
}

type worker struct {
	blockedTime int
	workingStep stp
}

func parse(input []string) []stp {
	var steps []stp
	for _, step := range input {
		r := regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin.`)
		strParsed := r.FindStringSubmatch(step)[1:]
		stepname := strParsed[0]
		before := strParsed[1]
		steps = append(steps, stp{name: stepname, before: before})
	}
	return steps
}

func canStepBeExecuted(tmp []stp, step string) bool {
	// step can be executed if step is
	for _, v := range tmp {
		if step == v.before {
			return false
		}
	}
	return true
}

func remove(a []stp, step stp) []stp {
	var tmp []stp
	for _, n := range a {
		if step.name != n.name {
			tmp = append(tmp, n)
		}
	}
	return tmp
}

func stepAlreadyInAvailable(stps []stp, s string) bool {
	// step can be executed if step is
	for _, v := range stps {
		if s == v.name {
			return true
		}
	}
	return false
}

func stepAlreadyInWorkers(step_ stp, wrkrs map[int]*worker) bool {
	// step can be executed if step is
	for _, v := range wrkrs {
		if step_.name == v.workingStep.name {
			return true
		}
	}
	return false
}

func initWorkers(n int) map[int]*worker {
	tmp := make(map[int]*worker)
	for i := 0; i < n; i++ {
		tmp[i] = &worker{}
	}
	return tmp
}

func main() {
	strInput := common.ReadToStringList("day7/input")
	parsed := parse(strInput)

	// part 1
	var steps []string
	for {
		var availableSteps []stp
		for _, el := range parsed {
			if canStepBeExecuted(parsed, el.name) {
				availableSteps = append(availableSteps, el)
			}
		}
		sort.Slice(availableSteps, func(i, j int) bool {
			return availableSteps[i].name < availableSteps[j].name
		})
		if len(availableSteps) > 1 {
			steps = append(steps, availableSteps[0].name)
			// remove used element from parsed
			parsed = remove(parsed, availableSteps[0])
		} else if len(availableSteps) == 1 {
			// when there is only one element left, we need to add step and it successor
			steps = append(steps, availableSteps[0].name)
			steps = append(steps, availableSteps[0].before)
			break
		}
	}
	fmt.Print("Part 1 result: ")
	for _, el := range steps {
		fmt.Printf("%s", el)
	}
	fmt.Println()

	// part 2
	// conf
	secondsPerStep := 60
	numOfWorkers := 5

	parsed = parse(strInput)
	steps = []string{}
	secondsSpent := 0
	workers := initWorkers(numOfWorkers)

	for {
		// get available tasks
		var availableSteps []stp
		for _, el := range parsed {
			if canStepBeExecuted(parsed, el.name) && !stepAlreadyInAvailable(availableSteps, el.name) {
				availableSteps = append(availableSteps, el)
			}
		}
		sort.Slice(availableSteps, func(i, j int) bool {
			return availableSteps[i].name < availableSteps[j].name
		})
		if len(parsed) == 1 && len(availableSteps) == 1 {
			// last line handle
			parsed = append(parsed, stp{availableSteps[0].before, ""})
		}
		if len(availableSteps) <= 0 {
			break
		}

		// give task to worker
		for _, step_ := range availableSteps {
			// give step_ to first available worker
			for _, wrkr := range workers {
				if !stepAlreadyInWorkers(step_, workers) {
					if wrkr.blockedTime == 0 {
						wrkr.workingStep = step_
						wrkr.blockedTime = secondsPerStep + int(step_.name[0]) - 64
					}
				}
			}
		}

		// reduce blocked time
		for _, wrkr := range workers {
			if wrkr.blockedTime != 0 {
				if wrkr.blockedTime == 1 {
					steps = append(steps, wrkr.workingStep.name)
					wrkr.blockedTime--
					parsed = remove(parsed, wrkr.workingStep)
				} else {
					wrkr.blockedTime--
				}
			}
		}

		secondsSpent++
	}
	fmt.Print("Part 2 result: ")
	for _, el := range steps {
		fmt.Printf("%s", el)
	}
	fmt.Println()
	fmt.Println("Part 2 seconds spent:", secondsSpent)

}
