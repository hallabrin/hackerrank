//convert the string to lower case letters, then byte array. then check if 97-122 are all present

package m

import "fmt"
import "bufio"
import "os"
import "io"
import "strings"

func main() {
	var s string
	read := bufio.NewReader(os.Stdin)
	s, err := read.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			fmt.Println(err.Error())
		}
	}

	s = strings.ToLower(s)
	bytes := []byte(s)
	var letters [26]bool

	for _, l := range bytes {
		if (l >= 97) && (l <= 122) {
			letters[l-97] = true
		}
	}

	for _, b := range letters {
		if b == false {
			fmt.Println("not pangram")
			return
		}
	}
	fmt.Println("pangram")
}
