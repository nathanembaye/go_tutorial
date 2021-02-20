package main

import "fmt"

func main() {
 ids :=[]int {123, 764, 32, 97}

 for i, id := range ids{
	 fmt.Printf("ID: %d  is --> %d\n", i, id)
 }

	sum := 0
 for _, id := range ids{
	sum += id
 }
 fmt.Println(sum)

 emails := map[string]string{"nathan":"nathan@gmail.com", "simon": "simon@gmail.com", "sharon": "sharon@gmail.com"}

 for i, j := range emails {
	 fmt.Printf("%s: %s\n", i, j)
 }

}