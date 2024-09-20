package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	if a>0{
		fmt.Printf("positive\n")
	} else if a<0{
		fmt.Printf("negative\n")
	} else {
		fmt.Printf("zero\n")
	}
	sum:=0
	for i:=1; i<11; i++{
		sum=sum+i

	}
	fmt.Println(sum)
	var day int
	fmt.Scanln(&day)
	switch day {
	case 1:
	  fmt.Println("Monday")
	case 2:
	  fmt.Println("Tuesday")
	case 3:
	  fmt.Println("Wednesday")
	case 4:
	  fmt.Println("Thursday")
	case 5:
	  fmt.Println("Friday")
	case 6:
	  fmt.Println("Saturday")
	case 7:
	  fmt.Println("Sunday")
	}
}

