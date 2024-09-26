package main

import (
	"fmt"
	"e.com/packages/fileops"
)

var accountBalanceFile string = "balanceFile.txt"

func main() {

	var bank_balance, err = fileops.ReadFileFloat(accountBalanceFile)
	if err != nil {
		fmt.Println()
		fmt.Println("error:")
		fmt.Println(err)
		fileops.WriteFileFloat(accountBalanceFile, 0)
		fmt.Println()
		fmt.Println("-- A new balance file has been created for you with 0 balance, but this session will be terminated... --")
		fmt.Println()
		panic("balance file not found, so terminating session")
	}

	// infinite loop to keep the code running until user exits with option 4
	for {

		giveOptions()

		fmt.Print("Enter your option : ")
		var userChoice int
		fmt.Scan(&userChoice)

		// perform tasks using if-else
		if userChoice == 1 {
			fmt.Println("Your bank balance is: ", bank_balance)
			fmt.Println("-----------------------------------------------------")

		} else if userChoice == 2 {
			fmt.Print("How much money you want to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Error: You cannot deposit negative or 0 value")
				fmt.Println("-----------------------------------------------------")
				continue
			}

			bank_balance += deposit
			fileops.WriteFileFloat(accountBalanceFile, bank_balance)

			fmt.Println("Your new bank balance is: ", bank_balance)
			fmt.Println("-----------------------------------------------------")

		} else if userChoice == 3 {
			fmt.Print("How much money you want to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw < 0 {
				fmt.Println("Error: You cannot enter a negative or 0 value")
				fmt.Println("-----------------------------------------------------")
				continue
			}

			if withdraw > bank_balance {
				fmt.Println("Sorry you are poor! Your account do not have this amount of money.")
				fmt.Println("-----------------------------------------------------")
				continue
			}

			bank_balance -= withdraw
			fileops.WriteFileFloat(accountBalanceFile, bank_balance)

			fmt.Println("Your new bank balance is: ", bank_balance)
			fmt.Println("-----------------------------------------------------")

		} else {
			fmt.Println("Thank you for banking with us, GoodBye!")
			fmt.Println("-----------------------------------------------------")
			fmt.Println()
			break
		}
	}
}
