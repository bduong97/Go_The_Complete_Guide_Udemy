package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultBalance = 1000.0
const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return defaultBalance, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultBalance, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance) //converting balance to string to write to file
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}
func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	for { //while loop
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye")
			return

		}
	}

}
