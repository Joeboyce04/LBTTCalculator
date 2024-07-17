package main

import (
	"errors"
	"fmt"
)

func TwoPercentCalculation(taxableAmount int)int {
	taxAmmountFloat:=float64(taxableAmount)

	return int((taxAmmountFloat/100)*2)
}

func FivePercentCalculation(taxableAmount int)int{
	taxAmmountFloat:=float64(taxableAmount)

	return int((taxAmmountFloat/100)*5)
}

func TenPercentCalculation(taxableAmount int)int{
	taxAmmountFloat:=float64(taxableAmount)

	return int((taxAmmountFloat/100)*10)
}

func TwelvePercentCalculation(taxableAmount int)int{
	taxAmmountFloat:=float64(taxableAmount)

	return int((taxAmmountFloat/100)*12)
}



func calculateLBTT(HousePrice int) (int , error) {
	if HousePrice<0{
		return 0, errors.New("cannot be negative houseprice") 
	}
	switch {
    case HousePrice <= 145000:
        return 0, nil
    case HousePrice <= 250000:
        return TwoPercentCalculation(HousePrice - 145000),nil
    case HousePrice <= 325000:
        return TwoPercentCalculation(250000 - 145000) + FivePercentCalculation(HousePrice - 250000), nil
	case HousePrice <= 750000:
		return TwoPercentCalculation(250000 - 145000) + FivePercentCalculation(325000 - 250000) + TenPercentCalculation(HousePrice-325000),nil
	default:
		return TwoPercentCalculation(250000 - 145000) + FivePercentCalculation(325000 - 250000) + TenPercentCalculation(750000-325000)+ TwelvePercentCalculation(HousePrice-750000),nil
	}
}

func HousePrice() int {
	HousePrice:=150000
	return HousePrice
}

func main() {
	HousePrice := HousePrice()
	lbttTax, _ := calculateLBTT(HousePrice)

	fmt.Println("This is the lbtt tax for a house of the value", HousePrice, "Tax:", lbttTax)
}