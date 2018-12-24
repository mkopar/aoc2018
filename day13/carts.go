package main

import (
	"aoc2018/lib/common"
	"fmt"
	"strings"
	"sort"
)

type cart struct {
	x, y             int
	orig, face, turn string
	crashed          bool
}

type cartsSort []*cart

func (crts cartsSort) Len() int {
	return len(crts)
}
func (crts cartsSort) Swap(i, j int) {
	crts[i], crts[j] = crts[j], crts[i]
}
func (crts cartsSort) Less(i, j int) bool {
	if crts[i].y == crts[j].y {
		return crts[i].x < crts[j].x
	} else {
		return crts[i].y < crts[j].y
	}
}

func getCarts(railsmap []string) []*cart {
	var tmp []*cart
	for y, line := range railsmap {
		for x, chr := range line {
			if string(chr) == "v" {
				tmp = append(tmp, &cart{x, y, "|", "down", "left", false})
			} else if string(chr) == "^" {
				tmp = append(tmp, &cart{x, y, "|", "up", "left", false})
			} else if string(chr) == "<" {
				tmp = append(tmp, &cart{x, y, "-", "left", "left", false})
			} else if string(chr) == ">" {
				tmp = append(tmp, &cart{x, y, "-", "right", "left", false})
			}
		}
	}
	return tmp
}

func printRails(in [][]string) {
	for _, line := range in {
		for _, chr := range line {
			fmt.Print(chr)
		}
		fmt.Println()
	}
}

// func checkCollisions - check if two carts have same x and y, if yes, return coords, otherwise -1 -1

func main() {
	in := common.ReadToStringList("day13/input")
	carts := getCarts(in)

	railsmap := make([][]string, len(in))
	for i := range railsmap {
		railsmap[i] = strings.Split(in[i], "")
	}

	crashx, crashy := 0, 0
	crashed := 0
	lastx, lasty := 0, 0
	firstx, firsty := 0, 0

	for {
		sort.Sort(cartsSort(carts))
		if crashed >= len(carts)-1 {
			fmt.Println(lastx, lasty)
			break
		}
		for _, crt := range carts {
			if crt.crashed {
				continue
			}
			if crashed == 2 {
				firstx, firsty = crashx, crashy
			}
			// move cart
			face := crt.face
			curx, cury := crt.x, crt.y
			if face == "up" {
				nextEl := string(railsmap[cury-1][curx])
				if nextEl == "<" || nextEl == ">" || nextEl == "v" || nextEl == "^" {
					crashx, crashy = curx, cury-1
					crt.crashed = true
					crashed++
					railsmap[cury][curx] = crt.orig
					for _, crt2 := range carts {
						// find cart on positions cury-1 curx
						if crt2.x == curx && crt2.y == cury-1 {
							crashed++
							crt2.crashed = true
							railsmap[crt2.y][crt2.x] = crt2.orig
						}
					}
				}
				if nextEl == "|" {
					railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, "^"
					crt.orig = "|"
					crt.y--
				} else if nextEl == "/" {
					railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, ">"
					crt.face = "right"
					crt.orig = "/"
					crt.y--
				} else if nextEl == "\\" {
					railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, "<"
					crt.face = "left"
					crt.orig = "\\"
					crt.y--
				} else if nextEl == "+" {
					if crt.turn == "left" {
						railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, "<"
						crt.face = "left"
						crt.turn = "straight"
					} else if crt.turn == "straight" {
						railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, "^"
						crt.face = "up"
						crt.turn = "right"
					} else if crt.turn == "right" {
						railsmap[cury][curx], railsmap[cury-1][curx] = crt.orig, ">"
						crt.face = "right"
						crt.turn = "left"
					}
					crt.orig = "+"
					crt.y--
				}
			} else if face == "down" {
				nextEl := string(railsmap[cury+1][curx])
				if nextEl == "<" || nextEl == ">" || nextEl == "v" || nextEl == "^" {
					crashx, crashy = curx, cury+1
					crt.crashed = true
					crashed++
					railsmap[cury][curx] = crt.orig
					for _, crt2 := range carts {
						if crt2.x == curx && crt2.y == cury+1 {
							crashed++
							crt2.crashed = true
							railsmap[crt2.y][crt2.x] = crt2.orig
						}
					}
				}
				if nextEl == "|" {
					railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, "v"
					crt.orig = "|"
					crt.y++
				} else if nextEl == "/" {
					railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, "<"
					crt.face = "left"
					crt.orig = "/"
					crt.y++
				} else if nextEl == "\\" {
					railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, ">"
					crt.face = "right"
					crt.orig = "\\"
					crt.y++
				} else if nextEl == "+" {
					if crt.turn == "left" {
						railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, ">"
						crt.face = "right"
						crt.turn = "straight"
					} else if crt.turn == "straight" {
						railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, "v"
						crt.face = "down"
						crt.turn = "right"
					} else if crt.turn == "right" {
						railsmap[cury][curx], railsmap[cury+1][curx] = crt.orig, "<"
						crt.face = "left"
						crt.turn = "left"
					}
					crt.orig = "+"
					crt.y++
				}
			} else if face == "left" {
				nextEl := string(railsmap[cury][curx-1])
				if nextEl == "<" || nextEl == ">" || nextEl == "v" || nextEl == "^" {
					crashx, crashy = curx-1, cury
					crt.crashed = true
					crashed++
					railsmap[cury][curx] = crt.orig
					for _, crt2 := range carts {
						if crt2.x == curx-1 && crt2.y == cury {
							crashed++
							crt2.crashed = true
							railsmap[crt2.y][crt2.x] = crt2.orig
						}
					}
				}
				if nextEl == "-" {
					railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "<"
					crt.orig = "-"
					crt.x--
				} else if nextEl == "/" {
					railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "v"
					crt.face = "down"
					crt.orig = "/"
					crt.x--
				} else if nextEl == "\\" {
					railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "^"
					crt.face = "up"
					crt.orig = "\\"
					crt.x--
				} else if nextEl == "+" {
					if crt.turn == "left" {
						railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "v"
						crt.face = "down"
						crt.turn = "straight"
					} else if crt.turn == "straight" {
						railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "<"
						crt.face = "left"
						crt.turn = "right"
					} else if crt.turn == "right" {
						railsmap[cury][curx], railsmap[cury][curx-1] = crt.orig, "^"
						crt.face = "up"
						crt.turn = "left"
					}
					crt.orig = "+"
					crt.x--
				}
			} else if face == "right" {
				nextEl := string(railsmap[cury][curx+1])
				if nextEl == "<" || nextEl == ">" || nextEl == "v" || nextEl == "^" {
					crashx, crashy = curx+1, cury
					crt.crashed = true
					crashed++
					railsmap[cury][curx] = crt.orig
					for _, crt2 := range carts {
						if crt2.x == curx+1 && crt2.y == cury {
							crashed++
							crt2.crashed = true
							railsmap[crt2.y][crt2.x] = crt2.orig
						}
					}
				}
				if nextEl == "-" {
					railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, ">"
					crt.orig = "-"
					crt.x++
				} else if nextEl == "/" {
					railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, "^"
					crt.face = "up"
					crt.orig = "/"
					crt.x++
				} else if nextEl == "\\" {
					railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, "v"
					crt.face = "down"
					crt.orig = "\\"
					crt.x++
				} else if nextEl == "+" {
					if crt.turn == "left" {
						railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, "^"
						crt.face = "up"
						crt.turn = "straight"
					} else if crt.turn == "straight" {
						railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, ">"
						crt.face = "right"
						crt.turn = "right"
					} else if crt.turn == "right" {
						railsmap[cury][curx], railsmap[cury][curx+1] = crt.orig, "v"
						crt.face = "down"
						crt.turn = "left"
					}
					crt.orig = "+"
					crt.x++
				}
			}
			left := 0
			for _, crt1 := range carts {
				if !crt1.crashed {
					left++
					lastx, lasty = crt1.x, crt1.y
				}
			}
			if left == 1 {
				fmt.Printf("Part 1: %d, %d\n", firstx, firsty)
				fmt.Printf("Part 2: %d, %d\n", lastx, lasty)
				return
			}
		}
		//printRails(railsmap)
	}
}
