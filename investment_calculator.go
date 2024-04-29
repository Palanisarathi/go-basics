package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var investmentAmount float64
	var years float64
	//var name string
	expectedReturnRate := 5.5
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	//fmt.Printf("investement %.1f", investmentAmount)
	//fmt.Print("Enter your name")
	//fmt.Scan(&name)
	//outputText(name)
	profit_calculator()

}

//	func outputText(text string) {
//		fmt.Println("hello" + text)
//	}
func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}
