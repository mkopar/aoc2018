package main

import (
	"aoc2018/lib/common"
	"fmt"
	"regexp"
	"strings"
	"math"
)

type point struct {
	px, py int
	vx, vy int
}

func parse(in []string) []point {
	var points []point
	for _, p := range in {
		// use only 1/10th points
		r := regexp.MustCompile(`position=<(.*)> velocity=<(.*)>`)
		strParsed := r.FindStringSubmatch(p)[1:]
		pos := strings.Split(strParsed[0], ",")
		vel := strings.Split(strParsed[1], ",")
		px := common.Atoi(strings.TrimSpace(pos[0]))
		py := common.Atoi(strings.TrimSpace(pos[1]))
		vx := common.Atoi(strings.TrimSpace(vel[0]))
		vy := common.Atoi(strings.TrimSpace(vel[1]))
		points = append(points, point{px: px, py: py, vx: vx, vy: vy})
	}
	return points
}

func main() {
	strInput := common.ReadToStringList("day10/input")
	initPoints := parse(strInput)
	// get bounding box
	diffx := math.MaxInt32
	diffy := math.MaxInt32
	s := 0
	maxx, maxy, minx, miny := 0, 0, 0, 0
	for i := 0; i < 20000; i++ {
		maxx, maxy, minx, miny = 0, 0, 0, 0
		for _, el := range initPoints {
			tmpx := el.px + i*el.vx
			tmpy := el.py + i*el.vy
			if tmpx > maxx {
				maxx = tmpx
			}
			if tmpx < minx {
				minx = tmpx
			}
			if tmpy > maxy {
				maxy = tmpy
			}
			if tmpy < miny {
				miny = tmpy
			}
		}
		s = i
		tmpdiffx := maxx - minx
		tmpdiffy := maxy - miny
		if tmpdiffx > diffx {
			break
		} else {
			diffx = tmpdiffx
		}
		if tmpdiffy > diffy {
			break
		} else {
			diffy = tmpdiffy
		}
	}

	s--
	fmt.Printf("bounding box %d x %d in second %d\n", diffx+1, diffy+1, s)

	grid := make([][]bool, diffy+1)
	for j := range grid {
		grid[j] = make([]bool, diffx+1)
	}

	for _, el := range initPoints {
		grid[el.py+el.vy*s][el.px+el.vx*s] = true
	}

	for _, row := range grid {
		tmp := ""
		for _, el := range row {
			if el {
				tmp += "#"
			} else {
				tmp += "."
			}
		}
		fmt.Println(tmp)
	}
}
