package main


func add [T int | float64 | string] (a, b T) T {
	return a+b
} //generic function, adapts at runtime to whatever is passed into it
