package main

import "fmt"

type tc struct {
	n, k int
	a    []int
}

func main() {
	var t int
	fmt.Scanln(&t)

	c := make([]tc, t)

	for i := range c {
		fmt.Scanf("%v %v", &c[i].n, &c[i].k)
		c[i].a = make([]int, c[i].n)
		for j := range c[i].a {
			fmt.Scanf("%v", &c[i].a[j])
		}
	}

	for i := range c {
		for _, j := range c[i].a {
			if j > 0 {
				c[i].n--
			}
		}
		if c[i].n < c[i].k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
