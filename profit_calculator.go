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

	// ebt := revenue - expense
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateProfit(revenue, expense, taxRate)
	fmt.Print("Earnings:")
	fmt.Printf("%.1f\n", ebt)
	fmt.Print("Profit:")
	fmt.Printf("%.1f\n", profit)
	fmt.Print("Ratio:")
	fmt.Printf("%.1f\n", ratio)
}
func calculateProfit(revenue float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
