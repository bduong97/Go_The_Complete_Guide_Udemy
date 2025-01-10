package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 
	var years float64 //Better readability 
	expectedReturnRate := 5.5 //When happy with inferred type, you can use a shorter way to assign a variable without var

	fmt.Print("Input your investment amount: ")
	fmt.Scan(&investmentAmount) //scan only accepts pointers
	//fmt.Scan(&inflationRate) //errors out because constants can't change
	
	fmt.Print("Input your years: ")
	fmt.Scan(&years) //scan only accepts pointers
	
	fmt.Print("Input your expected return rate: ")
	fmt.Scan(&expectedReturnRate) //scan only accepts pointers
	
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}