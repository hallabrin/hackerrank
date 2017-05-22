package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {

	var t int
	fmt.Scanln(&t)

	s := make([]string, t)
	read := bufio.NewReader(os.Stdin)
	for i := range s {
		s[i], _ = read.ReadString('\n')
		s[i] = strings.TrimSpace(s[i])
	}

	for _, l := range s {
		funny := true
		b := []byte(l)

		for i := 1; i < len(l); i++ {

			if abs((int(b[i])-97)-(int(b[i-1])-97)) != abs((int(b[len(b)-(i+1)])-97)-(int(b[len(b)-(i)])-97)) {
				funny = false
				//break
			}

		}
		if funny == false {
			fmt.Println("Not Funny")
		} else {
			fmt.Println("Funny")
		}

	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
