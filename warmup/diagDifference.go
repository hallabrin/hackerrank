package m

import "fmt"

func main() {
	size := 0
	fmt.Scanln(&size)

	numbers := make([]int, size*size)
	for i := 0; i < size*size; i++ {
		fmt.Scan(&numbers[i])
	}

	result := diagdiff(numbers, size)

	fmt.Println(result)
}

func diagdiff(numbers []int, size int) int {
	// primary diagonal numbers in the array are 0, (size+1), (size+1)*2, ...
	// secondary diagonal numbers is (size-1), (size-1)*2,, (size-1)*3,...

	sumP := 0
	sumS := 0
	for i := 0; i < size; i++ {
		posP := 0 + (size+1)*i
		posS := (size - 1) * (i + 1)
		sumP = sumP + numbers[posP]
		sumS = sumS + numbers[posS]
	}

	return abs(sumP - sumS)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
