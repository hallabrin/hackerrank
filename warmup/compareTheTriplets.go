package main

import (
	"fmt"
)

func main() {
	var a [3]int
	var b [3]int
	fmt.Scanf("%v %v %v\n", &a[0], &a[1], &a[2])
	fmt.Scanf("%v %v %v\n", &b[0], &b[1], &b[2])
	var sumA, sumB int
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			sumA++
		} else if a[i] < b[i] {
			sumB++
		}
	}
	fmt.Printf("%v %v", sumA, sumB)
}
