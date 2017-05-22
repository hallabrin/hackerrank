package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "io"

const bug = byte('X')

func main() {
	var n int
	fmt.Scanln(&n)

	s := make([]string, n)
	read := bufio.NewReader(os.Stdin)
	var err error
	for i, _ := range s {
		s[i], err = read.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
		}
		s[i] = strings.TrimSpace(s[i])
	}
	for i, row := range s {
		b := []byte(row)
		for j, v := range b {
			if v == bug {
				fmt.Printf("%d,%d", i, j)
			}
		}
	}

}
