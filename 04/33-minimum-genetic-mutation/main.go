package main

import "fmt"

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
}

func minMutation(start string, end string, bank []string) int {
	bank = append(bank, start)
	edges := map[string][]string{}
	for _, s := range bank {
		for _, s2 := range bank {
			d := 0
			for i := 0; i < 8; i++ {
				if s[i] != s2[i] {
					d++
				}
			}
			if d == 1 {
				edges[s] = append(edges[s], s2)
			}
		}
	}

	queue := []string{start, ""}
	seen := map[string]bool{}
	count := 0
	for len(queue) > 1 {
		s := queue[0]
		queue = queue[1:]
		if s == "" {
			queue = append(queue, "")
			count++
			continue
		}
		if s == end {
			return count
		}
		if seen[s] {
			continue
		}
		seen[s] = true

		queue = append(queue, edges[s]...)
	}

	return -1
}
