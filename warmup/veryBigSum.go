package m

import "fmt"

func main() {
	size := 0
	fmt.Scanln(&size)

	numbers := make([]uint64, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	result := veryBigSum(numbers)
	fmt.Println(result)
}

func veryBigSum(numbers []uint64) uint64 {
	var res uint64
	res = 0
	for _, n := range numbers {
		res = res + n
	}
	return res
}
