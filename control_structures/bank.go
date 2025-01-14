package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultBalance = 1000.0
const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultBalance, errors.New("failed to find file")
	}
	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultBalance, errors.New("failed to parse stored value")
	}
	return value, nil
}
func writeFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value) //converting balance to string to write to file
	os.WriteFile(fileName, []byte(balanceText), 0644)

}


func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	for { //while loop
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			fmt.Print("Withdrawl amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Withdrawal amount must be less than account balance")
				continue
			}
			accountBalance -= withdrawalAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye")
			return

		}
	}

}
