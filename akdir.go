package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		log.Print(e.Error())
		os.Exit(1)
	}
}

func main() {
	s := ""
	n := len(os.Args)
	if n < 2 && !terminal.IsTerminal(0) {
		fmt.Scan(&s)
	}
	m := utf8.RuneCountInString(s)
	if (m >= 1) || (m < 1 && n > 1) {
		if n > 1 {
			s = os.Args[1]
		}
	} else {
		check(errors.New("Error: not setted arg"))
	}
	err := os.MkdirAll(s, 0755)
	check(err)
}
