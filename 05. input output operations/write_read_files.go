package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalanceFile string = "balanceFile.txt"

// A funtion to create the "accountBalanceFile" if not present
// Or modify this file when withdraw | deposit happens
func writeToAccountFile(fileName string, balance float64) {
	// Sprint convert the float to string
	var modifiedBalance string = fmt.Sprint(balance)

	// []byte() can convert the string to []byte data type
	var modifiedBalanceInBytes []byte = []byte(modifiedBalance)

	// os.WriteFile() needs the data to be written to file
	// in []byte data-type first, then those []byte are written to file
	os.WriteFile(fileName, modifiedBalanceInBytes, 0644)
}

// A funtion to read the content of accountBalanceFile
// if the file is not present a error if will be returned.
func readAccountFile() (float64, error) {
	// os.ReadFile() return file's content as []byte
	balanceInByte, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0, errors.New("failed in finding the balance file")
	}
	// Converting the returned []byte back to string using string()
	balanceInString := string(balanceInByte)

	// Converting string back to float as bank balance written
	// to file was origanally in float64
	balanceInFloat, nil := strconv.ParseFloat(balanceInString, 64)
	if err != nil {
		return 0, errors.New("failed during parsing of the balance file")
	}

	return balanceInFloat, nil
}

func main() {

	var bank_balance, err = readAccountFile()
	if err != nil {
		fmt.Println()
		fmt.Println("error:")
		fmt.Println(err)
		writeToAccountFile(accountBalanceFile, 0)
		fmt.Println()
		fmt.Println("-- A new balance file has been created for you with 0 balance, but this session will be terminated... --")
		fmt.Println()
		panic("balance file not found, so terminating session")
	}

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
			writeToAccountFile(accountBalanceFile, bank_balance)

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
			writeToAccountFile(accountBalanceFile, bank_balance)

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
