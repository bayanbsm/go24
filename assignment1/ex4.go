package main

import "fmt"

func add (a, b int) int {
	return a + b
}

func swap (s,t string) (string, string) {
	return t,s
}

func division(n,m int)(int, int){
	q:= n/m
	r:=n-q*m
	return q,r
}
func main() {
	var s string
	var t string
	fmt.Scanln(&s, &t)
	var n int
	var m int
	fmt.Scanln(&n, &m)
	fmt.Println(add(5,6))
	fmt.Println(swap(s,t))
	fmt.Println(division(n,m))
}
