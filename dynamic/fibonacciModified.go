package main

import "fmt"
import "math/big"

var t []big.Int

func main() {
	var a, b big.Int
	var n int
	fmt.Scanf("%v %v %v", &a, &b, &n)

	t = make([]big.Int, n)

	t[0] = a
	t[1] = b

	for i := 2; i < n; i++ {
		t[i].Mul(&t[i-1], &t[i-1])
		t[i].Add(&t[i], &t[i-2])

	}

	fmt.Println(t[n-1].String())
}

// T(x) = T(x-1)² + T(x-2)
