package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name
}

func sum(num1, num2 int) int{
	return num1 + num2
}


func main() {
    fmt.Println(greeting("Nathan"))
	fmt.Println(sum(1, 2))
}