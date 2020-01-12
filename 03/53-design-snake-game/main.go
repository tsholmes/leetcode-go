package main

func main() {

}

type SnakeGame struct {
	next      map[[2]int][2]int
	wid, hei  int
	len, alen int
	head      [2]int
	tail      [2]int
	food      [][]int
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		next: map[[2]int][2]int{},
		wid:  width,
		hei:  height,
		len:  1,
		alen: 1,
		food: food,
	}
}

var dirs = map[string][2]int{
	"U": {0, -1},
	"L": {-1, 0},
	"D": {0, 1},
	"R": {1, 0},
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (s *SnakeGame) Move(direction string) int {
	d := dirs[direction]
	h2 := [2]int{s.head[0] + d[0], s.head[1] + d[1]}
	if h2[0] < 0 || h2[0] >= s.wid || h2[1] < 0 || h2[1] >= s.hei {
		return -1
	}
	s.next[s.head] = h2
	s.head = h2
	if _, ok := s.next[h2]; ok && s.head != s.tail {
		return -1
	}
	s.alen++
	if len(s.food) > 0 && s.food[0][1] == h2[0] && s.food[0][0] == h2[1] {
		s.len++
		s.food = s.food[1:]
	}
	if s.alen > s.len {
		ntail := s.next[s.tail]
		delete(s.next, s.tail)
		s.tail = ntail
		s.alen--
	}
	return s.len - 1
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
