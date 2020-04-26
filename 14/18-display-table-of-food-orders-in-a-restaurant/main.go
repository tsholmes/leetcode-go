package main

import (
	"sort"
	"strconv"
)

func main() {

}

func displayTable(orders [][]string) [][]string {
	tnums := map[string]bool{}
	items := map[string]bool{}
	itemCounts := map[[2]string]int{}

	for _, o := range orders {
		tnums[o[1]] = true
		items[o[2]] = true
		itemCounts[[2]string{o[1], o[2]}]++
	}

	var ilist []string
	for k := range items {
		ilist = append(ilist, k)
	}
	var tlist []string
	for k := range tnums {
		tlist = append(tlist, k)
	}

	sort.Strings(ilist)
	sort.Slice(tlist, func(i, j int) bool {
		a, _ := strconv.Atoi(tlist[i])
		b, _ := strconv.Atoi(tlist[j])
		return a < b
	})

	var res [][]string

	res = append(res, append([]string{"Table"}, ilist...))
	for _, t := range tlist {
		var row []string
		row = append(row, t)
		for _, i := range ilist {
			row = append(row, strconv.Itoa(itemCounts[[2]string{t, i}]))
		}
		res = append(res, row)
	}

	return res
}
