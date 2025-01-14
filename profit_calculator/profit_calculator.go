package main

import "fmt"
import "errors"
import "os"

// Goals
// 1) Validate user input
// 	=> show error message & exit if invalid input is provided
// 	- No negative numbers
// 	- Not 0
// 2) Store calculated results into file

func main() {

	revenue, err := getUserInput("Input revenue: ") 
	if err != nil { // multiple if statements if you want to output more specific error message depending on which input was incorrect
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Input expenses: ")
	if err != nil { // multiple if statements if you want to output more specific error message depending on which input was incorrect
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Input taxRate: ")
	if err != nil { // multiple if statements if you want to output more specific error message depending on which input was incorrect
		fmt.Println(err)
		return
	}

	//if err1 != nil || err2 != nil || err3 != nil { //useful if all the errors will be the same anyways; 
	//	fmt.Println(err1)
	//	return
	//}

	ebt, profit, ratio := finalCalcs(revenue, expenses, taxRate)
	resultsText := fmt.Sprintf("ebt: %.1f\nprofit: %.1f\nratio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("profit_calc.txt", []byte(resultsText), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var val float64
	fmt.Print(infoText)
	fmt.Scan(&val)
	if val <= 0 {
		return 0, errors.New("input must be must be greater than 0")
	}
	return val, nil
}
func finalCalcs(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio

}
