package main

import (
	"math"
	"math/rand"
)

func main() {

}

type Solution struct {
	radius, x_center, y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius:   radius,
		x_center: x_center,
		y_center: y_center,
	}
}

func (s *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64())
	a := 2.0 * math.Pi * rand.Float64()

	return []float64{
		s.x_center + r*s.radius*math.Cos(a),
		s.y_center + r*s.radius*math.Sin(a),
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
