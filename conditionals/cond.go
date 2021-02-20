package main

import "fmt"

func main() {

    x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y{
		fmt.Printf("%d is equals %d\n", y, x)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}


	
	color := "green"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither red or blue")
	}
	


}