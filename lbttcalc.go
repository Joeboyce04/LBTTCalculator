package main

import (
	"fmt"
)

func TwoPercentCalculation(HousePrice int)int {
	return (HousePrice/100)*2
}

func FivePercentCalculation(HousePrice int)int{
	return (HousePrice/100)*5
}

func TenPercentCalculation(HousePrice int)int{
	return (HousePrice/100)*10
}

func calculateLBTT(HousePrice int) int {
	switch {
    case HousePrice <= 145000:
        return 0
    case HousePrice <= 250000:
        return TwoPercentCalculation(HousePrice - 145000)
    case HousePrice <= 325000:
        return TwoPercentCalculation(250000 - 145000) + FivePercentCalculation(HousePrice - 250000)
    case HousePrice <= 750000:
        return TwoPercentCalculation(250000 - 145000) + FivePercentCalculation(325000 - 250000) + TenPercentCalculation(HousePrice - 325000)
	default:
		return 0
	}
}

func HousePrice() int {
	HousePrice:=150000
	return HousePrice
}

func main() {
	HousePrice := HousePrice()
	lbttTax := calculateLBTT(HousePrice)

	fmt.Println("This is the lbtt tax for a house of the value", HousePrice, "Tax:", lbttTax)
}