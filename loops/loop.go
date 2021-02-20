package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Fzzbuzz

	for i := 1; i <= 100; i++ {
		
		if i % 15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0{ 
			fmt.Println("Fizz")
		} else { 
			fmt.Println(i)
		}

	}

}