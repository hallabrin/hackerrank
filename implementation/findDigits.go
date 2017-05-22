package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	a := make([]int, t)

	for i := range a {
		fmt.Scanln(&a[i])
	}
	for _, n := range a {
		d := getDigits(n)
		c := 0
		for _, i := range d {
			if i == 0 {
				continue
			}
			if n%i == 0 {
				c = c + 1
			}
		}
		fmt.Println(c)
	}
}

func getDigits(x int) []int {
	var t []int
	for x >= 10 {

		t = append(t, (x % 10))
		x = x / 10
	}
	t = append(t, (x % 10))
	return t
}
