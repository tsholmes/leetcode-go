package main

func main() {

}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edges := map[string]map[string]float64{}
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := edges[a]; !ok {
			edges[a] = map[string]float64{}
		}
		if _, ok := edges[b]; !ok {
			edges[b] = map[string]float64{}
		}

		edges[a][b] = values[i]
		edges[b][a] = 1.0 / values[i]
	}

	var res []float64
	for _, q := range queries {
		if _, ok := edges[q[0]]; !ok {
			res = append(res, -1.0)
			continue
		}
		r := -1.0
		type state struct {
			n string
			v float64
		}
		work := []state{{q[0], 1.0}}
		seen := map[string]bool{}
		for len(work) > 0 {
			w := work[0]
			work = work[1:]
			if seen[w.n] {
				continue
			}
			seen[w.n] = true
			if w.n == q[1] {
				r = w.v
				break
			}
			for b, v := range edges[w.n] {
				work = append(work, state{b, w.v * v})
			}
		}
		res = append(res, r)
	}

	return res
}
