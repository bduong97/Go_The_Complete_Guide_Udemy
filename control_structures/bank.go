package main

import "fmt"

func main() {
	var accountBalance = 1000.0
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
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye")
			return

		}
	}

}
