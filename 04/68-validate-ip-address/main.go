package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("256.256.256.256"))
}

func validIPAddress(IP string) string {
	// ipv4
	{
		ps := strings.Split(IP, ".")
		if len(ps) == 4 {
			valid := true
			for _, p := range ps {
				v, err := strconv.Atoi(p)
				if err != nil {
					valid = false
					break
				}
				if v < 0 || v > 255 {
					valid = false
					break
				}
				if strconv.Itoa(v) != p {
					valid = false
					break
				}
			}
			if valid {
				return "IPv4"
			}
		}
	}
	// ipv6
	{
		ps := strings.Split(IP, ":")
		if len(ps) == 8 {
			valid := true
			for _, p := range ps {
				if len(p) == 0 || len(p) > 4 {
					valid = false
					break
				}
				if _, err := strconv.ParseUint(p, 16, 64); err != nil {
					valid = false
					break
				}
			}
			if valid {
				return "IPv6"
			}
		}
	}
	return "Neither"
}
