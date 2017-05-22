package m

import "fmt"

func main() {
	size := 0
	fmt.Scanln(&size)

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	result := simpleArraySum(numbers)
	fmt.Println(result)
}

func simpleArraySum(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res = res + n
	}
	return res
}
