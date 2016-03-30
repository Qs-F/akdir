package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
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
	s, err := homedir.Expand(s)
	check(err)
	err = os.MkdirAll(s, 0755)
	check(err)
	b := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(b, s)
	err = b.Flush()
	check(err)
}
