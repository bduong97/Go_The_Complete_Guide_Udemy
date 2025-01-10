package main

import "fmt"

func main () {
	var revenue, expenses, taxRate float64

	fmt.Print("Input revenue: ")
	fmt.Scan(&revenue)
	
	fmt.Print("Input expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input taxRate: ")
	fmt.Scan(&taxRate)

	ebt:=  revenue - expenses
	profit:= ebt * (1 - (taxRate/100)) 
	ratio := ebt/profit

	fmt.Println("ebt:", ebt)
	fmt.Println("profit:", profit)
	fmt.Println("ratio:", ratio)


}