package main

import "fmt"

func profit_calculator() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expense: ")
	fmt.Scan(&expense)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Print("Earnings:")
	fmt.Println(ebt)
	fmt.Print("Profit:")
	fmt.Println(profit)
	fmt.Print("Ratio:")
	fmt.Println(ratio)
}
