package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	m, n := len(s1), len(s2)
	i := 0
	for i < m && i < n {
		v1, _ := strconv.Atoi(s1[i])
		v2, _ := strconv.Atoi(s2[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
		i++
	}
	for i < m {
		v1, _ := strconv.Atoi(s1[i])
		if v1 > 0 {
			return 1
		}
		i++
	}
	for i < n {
		v2, _ := strconv.Atoi(s2[i])
		if v2 > 0 {
			return -1
		}
		i++
	}
	return 0
}

func main() {
	version1 := "1.01"
	version2 := "1.001"
	fmt.Println(compareVersion(version1, version2))
	version1 = "1.0"
	version2 = "1.0.0"
	fmt.Println(compareVersion(version1, version2))
	version1 = "0.1"
	version2 = "1.1"
	fmt.Println(compareVersion(version1, version2))
}
