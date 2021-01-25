package main

import "fmt"

type mongo map[int]string
type postgres map[int]string
type accessor interface {
	save(i int, s string)
	retrieve(i int) string
}

func (m mongo) save(i int, s string) {
	m[i] = s
}

func (m postgres) save(i int, s string) {
	m[i] = s
}

func (m mongo) retrieve(i int) string {
	return m[i]
}

func (m postgres) retrieve(i int) string {
	return m[i]
}

func put(a accessor, i int, s string) {
	a.save(i, s)
}

func get(a accessor, i int) string {
	return a.retrieve(i)
}

func main() {
	m := mongo{}
	p := postgres{}
	put(p, 1, "I'm a teapot")
	put(m, 1, "hello world")
	fmt.Println(get(m, 1))
	fmt.Println(get(p, 1))
}