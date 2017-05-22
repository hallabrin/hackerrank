package m

import "fmt"

func main() {
	size := 0
	fmt.Scanln(&size)

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	result := arrayDS(numbers)

	fmt.Println(result)
}

func arrayDS(numbers []int) string {
	var res string
	for i := len(numbers); i > 0; i-- {
		res = res + fmt.Sprint(numbers[i-1]) + " "
	}
	return res
}
