package main

import (
	"fmt"
)

func main() {
	var n [5]int64
	fmt.Scanf("%v %v %v %v %v\n", &n[0], &n[1], &n[2], &n[3], &n[4])
	// find lowest number
	high := int64(0)
	low := int64(9223372036854775807)
	var sum int64
	for i, _ := range n {
		if n[i] > high {
			high = n[i]
		}
		if n[i] < low {
			low = n[i]
		}
		sum += n[i]
	}
	fmt.Printf("%v %v", sum-high, sum-low)
}
