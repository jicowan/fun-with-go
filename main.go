package main

import "fmt"

type foo struct {
	word string 
}

type foobar interface {
	bar()
}

func main() {
	fmt.Println("Hello world!")
	f := foo {
		word: "jkjkjkj",
	}
	f.bar()
	submain(&f)
}

func submain(f foobar) {
	f.bar()
}

func (f *foo) bar() {
	fmt.Println("Hello from foo", f.word)
}