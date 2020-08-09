package main

import (
	"path"
	"strings"
)

func main() {

}

func findDuplicate(paths []string) [][]string {
	groups := map[string][]string{}

	for _, line := range paths {
		ps := strings.Split(line, " ")
		dir := ps[0]
		for _, f := range ps[1:] {
			pidx := strings.IndexByte(f, '(')
			name := f[:pidx]
			content := f[pidx+1 : len(f)-1]

			groups[content] = append(groups[content], path.Join(dir, name))
		}
	}

	var res [][]string
	for _, g := range groups {
		if len(g) > 1 {
			res = append(res, g)
		}
	}

	return res
}
