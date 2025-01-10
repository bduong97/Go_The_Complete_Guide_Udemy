package main

import "fmt"

func main () {
	var revenue, expenses, taxRate float64

	revenue = getUserInput("Input revenue: ")
	expenses = getUserInput("Input expenses: ")
	taxRate = getUserInput("Input taxRate: ")

	ebt, profit, ratio := finalCalcs(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var val float64
	fmt.Print(infoText)
	fmt.Scan(&val)
	return val
}
func finalCalcs(revenue, expenses, taxRate float64 ) (float64, float64, float64) {
	ebt:=  revenue - expenses
	profit:= ebt * (1 - (taxRate/100)) 
	ratio := ebt/profit
	return ebt, profit, ratio

}