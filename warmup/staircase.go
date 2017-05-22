package m

import "fmt"
import "strings"

func main() {
	height := 0
	fmt.Scanln(&height)

	for i := 0; i < height; i++ {
		l := strings.Repeat(" ", height-i-1) + strings.Repeat("#", i+1)
		fmt.Println(l)
	}
}
