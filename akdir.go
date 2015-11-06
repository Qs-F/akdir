package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Print("Error:not setted arg")
		os.Exit(1)
	}
	err := os.MkdirAll(os.Args[1], 0755)
	check(err)
}
