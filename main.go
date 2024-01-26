package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b strings.Builder
	s := "d\n5abc"
	for i := 0; i < len(s); i++ {
		if s[i] == 10 {
			b.WriteString(string(s[i]))
			fmt.Println("sd ", b.String())
		}
		if s[i] <= 57 && s[i] != 10 {
			if i == 0 {
				fmt.Println("error")
				break
			}
			if i < len(s)-1 && s[i+1] <= 57 {
				fmt.Println("error2")
				break
			}
		}
		if i < len(s)-2 && s[i+1] <= 57 && s[i+2] > 57 {
			repeat, _ := strconv.Atoi(string(s[i+1]))
			b.WriteString(strings.Repeat(string(s[i]), repeat))
			i++
			continue
		}

		fmt.Println(s[i], i)
		b.WriteString(string(s[i]))
	}

	fmt.Println(b.String())
}
