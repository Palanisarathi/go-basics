package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice is: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice:", choice)
		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Println("Enter the amount to deposit:")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid deposite amount. Must be greater than 0")
				continue
			}

			accountBalance += deposit
			fmt.Print("Your new balance is ", accountBalance)
		} else if choice == 3 {
			fmt.Println("Enter the amount to withdraw:")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid amount. Amount must be greater than 0")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("withdrawal amount is greater than accountbalance")
				continue
			}
			accountBalance -= withdraw
			fmt.Print("Your new balance is ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

}

//const inflationRate float64 = 2.5
// var investmentAmount float64
// var years float64
// //var name string
// expectedReturnRate := 5.5
// fmt.Print("Investment Amount: ")
// fmt.Scan(&investmentAmount)
// fmt.Print("Years: ")
// fmt.Scan(&years)
// fmt.Print("Expected Return Rate: ")
// fmt.Scan(&expectedReturnRate)
// futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
// //futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// //futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
// fmt.Println(futureValue)
// fmt.Println(futureRealValue)
// //fmt.Printf("investement %.1f", investmentAmount)
// //fmt.Print("Enter your name")
// //fmt.Scan(&name)
// //outputText(name)
// profit_calculator()

//	func outputText(text string) {
//		fmt.Println("hello" + text)
//	}
// func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, frv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	frv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
