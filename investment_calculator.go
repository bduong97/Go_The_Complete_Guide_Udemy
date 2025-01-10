package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000 //Multiple variable declaration
	var years float64 = 10 //Better readability 
	expectedReturnRate := 5.5 //When happy with inferred type, you can use a shorter way to assign a variable without var

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	fmt.Println(futureValue)

}