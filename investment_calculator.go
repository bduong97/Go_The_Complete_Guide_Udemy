package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64 
	var years float64 //Better readability 
	expectedReturnRate := 5.5 //When happy with inferred type, you can use a shorter way to assign a variable without var

	//fmt.Print()
	outputText("Input your investment amount: ")
	
	fmt.Scan(&investmentAmount) //scan only accepts pointers
	//fmt.Scan(&inflationRate) //errors out because constants can't change
	
	//fmt.Print()
	outputText("Input your expected return rate: ")
	fmt.Scan(&expectedReturnRate) //scan only accepts pointers

	//fmt.Print()
	outputText("Input your years: ")
	fmt.Scan(&years) //scan only accepts pointers
	
	
	futureValue, futureRealValue :=	calculatorFutureValue(investmentAmount, expectedReturnRate, years)
	
	formattedFV := 	fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)
	//fmt.Println(futureValue)
	//fmt.Printf("Future Value: %.1f\n", futureValue)
	//fmt.Printf("Future real value: %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculatorFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv //alternative way to return values
}