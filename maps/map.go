package main

import "fmt"

func main() {
/*
 emails := make(map[string]string)
 emails["nathan"] = "nathan@gmail.com"
 emails["simon"] = "simon@gmail.com"
 emails["sharon"] = "sharon@gmail.com"
*/
 
 emails := map[string]string{"nathan":"nathan@gmail.com", "simon": "simon@gmail.com", "sharon": "sharon@gmail.com"}
 fmt.Println(emails)
 fmt.Println(len(emails))
 fmt.Println(emails["nathan"])

 delete(emails, "simon")
 fmt.Println(emails)
}