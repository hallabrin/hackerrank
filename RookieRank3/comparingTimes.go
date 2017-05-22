package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "io"
import "time"

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
				//	fmt.Println(err.Error())
			}
		}
		s[i] = strings.TrimSpace(s[i])
	}
	for _, line := range s {
		times := strings.Split(line, " ")
		if len(times) < 2 {
			//	fmt.Println("could not find 2 times")
		}
		t1, err := time.Parse(time.Kitchen, times[0])
		if err != nil {
			//	fmt.Println("could not parse first time:", times[0], err.Error())
		}
		t2, err := time.Parse(time.Kitchen, times[1])
		if err != nil {
			//	fmt.Println("could not parse second time:", times[1], err.Error())
		}
		// no need to check if t1=t2
		if t1.Sub(t2) < 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}
