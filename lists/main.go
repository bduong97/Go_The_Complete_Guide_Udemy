package main

import "fmt"

type floatMap map[string]float64 //type alias

func main() {
	userNames := make([]string, 2, 5) //type, length, capacity
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manual")
	fmt.Println(userNames)

	courseRatings := make (floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["java"] = 4.5 //reallocating memory since exceeding map capacity of 3

	for index, value := range userNames { //for loop when going through away
		fmt.Println("Index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRatings { //same syntax for looping through maps
		fmt.Println(key)
		fmt.Println(value)
	}
}