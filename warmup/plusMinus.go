package m

import "fmt"

func main() {
	size := 0
	fmt.Scanln(&size)

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	p, n, z := plusMinus(numbers)

	fmt.Print(p, "\n", n, "\n", z)
}

func plusMinus(numbers []int) (float32, float32, float32) {
	// primary diagonal numbers in the array are 0, (size+1), (size+1)*2, ...
	// secondary diagonal numbers is (size-1), (size-1)*2,, (size-1)*3,...
	var p  float32 
	var n  float32 
    var z  float32 
	l := float32(len(numbers))
	for _, n := range numbers {
		switch  {
		case n > 0:
			p = p + 1
		case n < 0:
			n = n + 1
		case n == 0:
			z = z + 1
		}
	}

	return p / l, n / l, z / l
}
