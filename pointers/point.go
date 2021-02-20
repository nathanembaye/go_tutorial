package main

import "fmt"

func main() {
		a := 5
		b := &a

		fmt.Println(a,b)
		fmt.Printf("%T\n", a)

		//Value at address 
		fmt.Println(*b)
		fmt.Println(*&a)

		//Change value via a pointer
		*b = 10
		fmt.Println(a)
}