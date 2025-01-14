package main

import "fmt"

func main() {
	age := 32 //regular variable
	//var agePointer *int = &age
	agePointer := &age //& is used to get memory address
	fmt.Println("Age:", *agePointer) //* is used to get the value stored at memory address (dereferencing)
	getAdultYears(agePointer)
	fmt.Println(*agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}