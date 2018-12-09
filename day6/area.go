package main

import (
	"strings"
	"aoc2018/lib/common"
	"math"
	"fmt"
)

type coordinate struct {
	x int
	y int
}

type point struct {
	coord coordinate
	id    int
}

func (p1 point) distance(x int, y int) float64 {
	return math.Abs(float64(p1.coord.x-x)) + math.Abs(float64(p1.coord.y-y))
}

func parse(input []string) ([]point, int, int) {
	var tmp []point
	maxx, maxy := 0, 0

	for idx, el := range input {
		tmpStr := strings.Split(el, ", ")
		xInt := common.Atoi(tmpStr[0])
		yInt := common.Atoi(tmpStr[1])
		tmp = append(tmp, point{id: idx + 1, coord: coordinate{x: xInt, y: yInt}})
		if xInt > maxx {
			maxx = xInt
		}
		if yInt > maxy {
			maxy = yInt
		}
	}

	return tmp, maxx, maxy
}

func main() {
	strInput := common.ReadToStringList("day6/input")
	parsedPoints, maxx, maxy := parse(strInput)

	grid := make(map[coordinate]int) // {"coordinate(x,y)": "id"}

	// part 2
	gridRegion := make(map[coordinate]int) // same as grid, just for part two
	regionLimit := 10000

	// populate grid with ids
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			minId := -1
			minDist := math.MaxFloat64
			regionDist := 0.0
			for _, pnt := range parsedPoints {
				dist := pnt.distance(x, y)
				if dist < minDist {
					minId = pnt.id
					minDist = dist
				} else if dist == minDist {
					minId = 0
				}
				// part 2
				regionDist += dist
			}
			grid[coordinate{y: y, x: x}] = minId

			// part 2
			if regionDist < float64(regionLimit) {
				gridRegion[coordinate{y: y, x: x}] = 1
			}
			//fmt.Printf("%2d ", minId)
		}
		//fmt.Println()
	}

	edgePointIds := []int{0}
	areaSizes := make(map[int]int) // {"id": "area_size"}
	// part 2
	regionSize := 0

	// get edge ids and calculate areas
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			id := grid[coordinate{x: x, y: y}]
			if y == 0 || x == 0 || y == maxy || x == maxx {
				if !common.IntListContains(edgePointIds, id) {
					edgePointIds = append(edgePointIds, id)
				}
			} else {
				areaSizes[id]++
			}

			// part 2
			if gridRegion[coordinate{x: x, y: y}] == 1 {
				regionSize++
			}
		}
	}

	// get max areaSizes
	maxArea := 0
	maxAreaId := -1
	for k, v := range areaSizes {
		if !common.IntListContains(edgePointIds, k) {
			if v > maxArea {
				maxArea = v
				maxAreaId = k
			}
		}
	}
	fmt.Printf("Part 1 max areaSizes id: %d, areaSizes: %d\n", maxAreaId, maxArea)
	fmt.Printf("Part 2 region size: %d\n", regionSize)

}
