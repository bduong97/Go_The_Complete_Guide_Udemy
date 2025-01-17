package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) //functions as parameters
	tripled := transformNumbers(&numbers, triple) //functions as parameters
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{} //empty dynamic array of ints
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
