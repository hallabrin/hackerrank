package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	c := make([]int, t)

	for i := range c {
		fmt.Scanln(&c[i])
	}
	for _, n := range c {
		h := 1
		for i := 0; i < n; i++ {
			switch i % 2 {
			case 0:
				h = h * 2
			case 1:
				h = h + 1
			}
		}
		fmt.Println(h)
	}

}
