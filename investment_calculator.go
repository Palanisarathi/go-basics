package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	years, expectedReturnRate := 10.0, 5.5
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
