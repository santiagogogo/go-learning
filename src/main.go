package main

import "fmt"

const (
	a=iota
	b
	c
	d
	e
)

func main()  {
	fmt.Println("hello world")
	fmt.Println(sum(1,2))
	fmt.Println(sum(d,e))


}

func sum(a,b int) int {
	return a+b
}