package main

import (
	"errors"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 1 {
		check(errors.New("not setted arg"))
	}
	err := os.MkdirAll(os.Args[1], 0755)
	check(err)
}
