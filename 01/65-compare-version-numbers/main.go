package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
}

func compareVersion(version1 string, version2 string) int {
	v1ps := strings.Split(version1, ".")
	v2ps := strings.Split(version2, ".")
	l := len(v1ps)
	if len(v2ps) > l {
		l = len(v2ps)
	}
	v1 := make([]int, l)
	for i, v := range v1ps {
		v1[i], _ = strconv.Atoi(v)
	}
	v2 := make([]int, l)
	for i, v := range v2ps {
		v2[i], _ = strconv.Atoi(v)
	}
	for i := range v1 {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}
	return 0
}
