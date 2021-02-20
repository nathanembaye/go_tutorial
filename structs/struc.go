package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) greet() string {
	return "Hello my name is: " + p.firstName + " " +  p.lastName +" and im " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthdy() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	
	if p.gender == "M" {
		return
	} else{
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person 	{firstName: "Nathan", lastName: "embaye", city: "Ottawa", gender: "M", age: 23}

	person2 := Person 	{"GirlFirstName", "GirlLastName", "Toronto", "F", 23}

	fmt.Println(person1)
	person1.hasBirthdy()
	person1.getMarried("GirlLastName")
	person2.getMarried("embaye")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	
}