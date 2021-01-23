package main

import "fmt"

type foo struct {
	word string 
}

func main() {
	fmt.Println("Hello world!")
	f := foo {
		word: "jkjkjkj",
	}
	f.bar()
}

func (f *foo) bar() {
	fmt.Println("Hello from foo", f.word)
}