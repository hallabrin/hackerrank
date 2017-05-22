package main

import "fmt"

func main() {
	v := 0
	fmt.Scanln(&v)
	n := 0
	fmt.Scanln(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	result := findPosition(array, v)

	fmt.Println(result)
}

func findPosition(numbers []int, v int) int {
	for i, n := range numbers {
		if n == v {
			return i
		}
	}
	return -1
}
