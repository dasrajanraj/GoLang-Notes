package main

import "fmt"

//Declaring a variable at package level
/*
Below example if, variable is used internally in package.
*/
var name string = "Golang"

/*
Variable exported across package; Capitalize
*/
var NAME string = "http://www.google.com"

func main() {

	var (
		firstName   string = "Rajan"
		address     string = "ABC"
		company     string = "XYZ"
		phoneNumber int    = 123
	)
	fmt.Println(firstName, address, company, phoneNumber)

	var name string = "Golang"
	fmt.Println(name) // Golang
	//Printing a type of the variable
	fmt.Printf("%T", name) // string

	var isCorrect bool = true
	if isCorrect {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
