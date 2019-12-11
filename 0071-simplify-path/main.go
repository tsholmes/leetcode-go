package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//./"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

func simplifyPath(path string) string {
	ps := strings.Split(path[1:], "/")
	count := 0
	for _, p := range ps {
		switch p {
		case "":
			// noop
		case ".":
			// noop
		case "..":
			if count > 0 {
				count--
			}
		default:
			ps[count] = p
			count++
		}
	}
	return "/" + strings.Join(ps[:count], "/")
}
