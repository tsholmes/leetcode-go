package main

import "fmt"

func main() {
	r := Robot{pos: [2]int{1, 3}}
	cleanRoom(&r)
}

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

var (
	rdirs = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	board = [5][8]int{
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
)

type Robot struct {
	pos [2]int
	dir int
}

func (r *Robot) Move() bool {
	p2 := [2]int{r.pos[0] + rdirs[r.dir][0], r.pos[1] + rdirs[r.dir][1]}
	if p2[0] < 0 || p2[0] >= 5 || p2[1] < 0 || p2[1] >= 8 || board[p2[0]][p2[1]] == 0 {
		return false
	}
	r.pos = p2
	return true
}
func (r *Robot) TurnLeft()  { r.dir = (r.dir + 3) & 3 }
func (r *Robot) TurnRight() { r.dir = (r.dir + 1) & 3 }
func (r *Robot) Clean()     { fmt.Println(r.pos) }

func cleanRoom(robot *Robot) {
	poss := map[[2]int]bool{}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var pos [2]int
	var dir int

	poss[pos] = true
	robot.Clean()

	search := func(start [2]int) ([][2]int, bool) {
		queue := [][2][2]int{{start, start}}
		back := map[[2]int][2]int{}
		var end [2]int
		var found bool
		for len(queue) > 0 {
			pp := queue[0]
			queue = queue[1:]
			if _, ok := back[pp[0]]; ok {
				continue
			}
			back[pp[0]] = pp[1]
			if open, ok := poss[pp[0]]; !ok {
				end = pp[0]
				found = true
				break
			} else if !open {
				continue
			}

			for _, d := range dirs {
				p2 := [2]int{pp[0][0] + d[0], pp[0][1] + d[1]}
				queue = append(queue, [2][2]int{p2, pp[0]})
			}
		}
		if !found {
			return nil, false
		}
		var path [][2]int
		for end != start {
			path = append(path, end)
			end = back[end]
		}
		for i := 0; i < len(path)/2; i++ {
			j := len(path) - i - 1
			path[i], path[j] = path[j], path[i]
		}
		return path, true
	}

	follow := func(path [][2]int) bool {
		for i, p := range path {
			d := [2]int{p[0] - pos[0], p[1] - pos[1]}
			var ndir int
			switch d {
			case dirs[0]:
				ndir = 0
			case dirs[1]:
				ndir = 1
			case dirs[2]:
				ndir = 2
			case dirs[3]:
				ndir = 3
			default:
				panic(fmt.Sprint("invalid hop", pos, p))
			}
			for dir != ndir {
				robot.TurnRight()
				dir = (dir + 1) & 3
			}
			ok := robot.Move()
			if !ok {
				if i == len(path)-1 {
					return false
				}
				panic(fmt.Sprint("invalid path", path))
			}
			pos = p
		}
		return true
	}

	for {
		path, ok := search(pos)
		if !ok {
			break
		}
		pend := path[len(path)-1]
		poss[pend] = follow(path)
		if poss[pend] {
			robot.Clean()
		}
	}
}
