package main

import "fmt"

// 5, 10,20 每一杯5
func lemonadeChange(bills []int) bool {
	var count [3]int
	for _, b := range bills {
		if b == 5 {
			count[0]++
		} else if b == 10 {
			if count[0] <= 0 {
				return false
			}
			count[0]--
			count[1]++
		} else {
			if count[1] > 0 && count[0] > 0 {
				count[0]--
				count[1]--
				count[2]++
			} else if count[0] >= 3 {
				count[0] -= 3
				count[2]++
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	bills := []int{5, 5, 5, 10, 20}
	fmt.Println(lemonadeChange(bills))
	bills = []int{5, 5, 10, 10, 20}
	fmt.Println(lemonadeChange(bills))
}
