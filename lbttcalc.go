package main

import (
	"fmt"
)

func TwoPercentCalculation(HousePrice int)int {
	return (HousePrice/100)*2
}

func calculateLBTT(HousePrice int) int {
	switch {
	case HousePrice <= 145000:
		return 0
	case HousePrice <= 250000:
		return  TwoPercentCalculation(HousePrice- 145000)

	default:
		return 0
	}
}

func HousePrice() int {
	return 150000
}

func main() {
	HousePrice := HousePrice()
	lbttTax := calculateLBTT(HousePrice)

	fmt.Println("This is the lbtt tax for a house of the value", HousePrice, "Tax:", lbttTax)
}