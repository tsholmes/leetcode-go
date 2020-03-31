package main

func main() {

}

type state struct {
	inName string
	t      int
}

type UndergroundSystem struct {
	inFlight map[int]state
	sums     map[[2]string]int
	counts   map[[2]string]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		inFlight: map[int]state{},
		sums:     map[[2]string]int{},
		counts:   map[[2]string]int{},
	}
}

func (u *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	u.inFlight[id] = state{
		inName: stationName,
		t:      t,
	}
}

func (u *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	prev := u.inFlight[id]
	delete(u.inFlight, id)
	dur := t - prev.t
	k := [2]string{prev.inName, stationName}
	u.sums[k] += dur
	u.counts[k]++
}

func (u *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	k := [2]string{startStation, endStation}
	return float64(u.sums[k]) / float64(u.counts[k])
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
