package main

import (
	"fmt"
)

func main() {
	size := 0
	fmt.Scanln(&size)

	c := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&c[i])
	}

	var high int //hight of the heighest candle
	var sum int  //numbers of candles with she blows out

	for _, n := range c {
		if n > high {
			high = n
			sum = 1
		} else if n == high {
			sum++
		}
	}
	fmt.Println(sum)

}
