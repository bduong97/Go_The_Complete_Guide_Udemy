package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int { //useful if you want to allow either slice of ints or ints individually
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
