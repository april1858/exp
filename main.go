package main

import (
	"errors"
	"fmt"
	"strings"
)

var InvalidInput = errors.New("invalid input")

func Unpack(strIn string) (string, error) {
	strOut := ""
	var buff string
	fmt.Printf("%s ", string(rune(10)))
	for _, c := range strIn {
		if c == 10 {
			c = 92
		}
		fmt.Println("cc - ", c)
		if c <= 57 {
			if buff == "" {
				return "", InvalidInput
			} else {
				strOut += strings.Repeat(buff, int(c-'0'))
				buff = ""
			}
		}
		if c > 57 {

			if buff == "" {
				if c == 92 {
					buff = string(c) + "n"
				} else {
					buff = string(c)
				}

			} else {
				if c == 92 {
					strOut += buff
					buff = string(c) + "n"
				} else {
					strOut += buff
					buff = string(c)
				}

			}
		}
	}
	if buff != "" {
		strOut += string(buff)
	}
	return strOut, nil
}

func main() {
	s, err := Unpack("a\n2b3nc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
