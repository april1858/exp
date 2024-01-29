package main

import (
	"errors"
	"fmt"
	"strings"
)

var InvalidInput = errors.New("invalid input")

func Unpack(strIn string) (string, error) {
	strOut := ""
	var buff, s string
	slash := false
	for _, c := range strIn {
		if c == 10 {
			slash = true
		}
		if c <= 53 && c != 10 {
			if buff == "" {
				return "", InvalidInput
			} else {
				strOut += strings.Repeat(buff, int(c-'0'))
				buff = ""
			}
		}
		if c > 53 {

			if buff == "" {
				if slash == true {
					s = string(rune(10)) + string(c)
					buff = s
					fmt.Println("s ", s)
				} else {
					buff = string(c)
				}

			} else {
				strOut += buff
				buff = string(c)
			}
		}
	}
	if buff != "" {
		strOut += string(buff)
	}
	return strOut, nil
}

func main() {
	s, err := Unpack("d\n5abc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
