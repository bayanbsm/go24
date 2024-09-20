package main

import "fmt"

func main() {
	var a int = 12
	var b string = "simba"
	c := 1.65
	var d bool = false
	fmt.Printf("%d %T\n",a, a )
	fmt.Printf("%s %T\n",b, b )
	fmt.Printf("%f %T\n",c, c )
	fmt.Printf("%b %T\n",d, d )
}