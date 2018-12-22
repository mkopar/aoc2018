package main

import "fmt"

func topleftNxN(grid [300][300]int, squareLen int) (int, int, int) {
	maxx := 0
	maxy := 0
	sum := 0
	for x := 0; x < len(grid)-squareLen; x++ {
		for y := 0; y < len(grid[0])-squareLen; y++ {
			tmpSum := 0
			for i := 0; i < squareLen; i++ {
				for j := 0; j < squareLen; j++ {
					tmpSum += grid[x+i][y+j]
				}
			}
			if tmpSum > sum {
				sum = tmpSum
				maxx = x
				maxy = y
			}
		}
	}
	return maxx, maxy, sum
}

func main() {
	grid := [300][300]int{}
	serial := 9995 // example --> input
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += serial
			powerLevel *= rackId
			powerLevel = int(powerLevel/100) % 10 // keep hundreds only
			powerLevel -= 5
			grid[x][y] = powerLevel
		}
	}
	x, y, _ := topleftNxN(grid, 3)
	fmt.Printf("Part 1: %d,%d\n", x, y)
	var maxx, maxy, maxSum, tmpSum, maxSq int
	for i := 0; i < len(grid); i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
		x, y, tmpSum = topleftNxN(grid, i)
		if tmpSum > maxSum {
			maxSum = tmpSum
			maxSq = i
			maxx = x
			maxy = y
		}
	}

	fmt.Printf("Part 2: %d,%d,%d\n", maxx, maxy, maxSq)
}
