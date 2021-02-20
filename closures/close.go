package main

import "fmt"

//adder() returns an anonymous function, which has int params and returns int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	sum := adder()

	for i := 0; i <10; i++ {
		fmt.Println(sum(i))
	}
	 
}