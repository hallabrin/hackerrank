package m

import "fmt"
import "time"

func main() {
	var s string
	fmt.Scanln(&s)

	t,_ := time.Parse("3:04:05PM",s)
    fmt.Println(t.Format("15:04:05"))
}
