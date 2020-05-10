package main

func main() {

}

func destCity(paths [][]string) string {
	all := map[string]bool{}
	out := map[string]bool{}

	for _, p := range paths {
		all[p[0]] = true
		all[p[1]] = true
		out[p[0]] = true
	}

	for k := range all {
		if !out[k] {
			return k
		}
	}

	return ""
}
