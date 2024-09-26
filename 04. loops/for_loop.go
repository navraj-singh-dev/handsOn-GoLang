package main

import "fmt"

var bank_balance float64 = 1000

func main() {

	fmt.Println()
	fmt.Println("Welcome to Navraj Bank!")

	// simple iterative for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("I am a simple for loop and run only a limited amount of iterations")
	// }

	// infinite loop to keep the code running until user exits with option 4
	for {
		// greet and give options and store user input
		fmt.Println()
		fmt.Println("What Do You Want To Do ?")
		fmt.Println()
		fmt.Println("1. Check Your Bank Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

		fmt.Print("Enter your option : ")
		var userChoice int
		fmt.Scan(&userChoice)

		fmt.Println()
		fmt.Println("Your Chose Option", userChoice)
		fmt.Println()

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
