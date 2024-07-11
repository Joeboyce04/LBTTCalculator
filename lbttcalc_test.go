package mymain

import (
	"testing"
)

func TestLBTTBelowFirstValue(t *testing.T){
	HousePrice:= 100000

	got:=calculateLBTT(HousePrice)

	want:=0

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsFirstVlue(t *testing.T){
	HousePrice:=145000

	got:=calculateLBTT(HousePrice)

	want:=0

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}


func TestLBTTBelowSecondValue(t *testing.T){
	HousePrice:= 180000

	got:= calculateLBTT(HousePrice)

	want:=3600

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsSecondValue(t *testing.T){
	HousePrice:= 250000

	got:= calculateLBTT(HousePrice)

	want:=5000

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}


func TestTwoPercentCalculation(t *testing.T){
	HousePrice:=150000

	got:=TwoPercentCalculation(HousePrice)

	want:=3000

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}

} 

func TestHousePriceFunction(t *testing.T){
	 got := HousePrice()

	 want := 150000 

	 if got != want {
		 t.Error("Got", got, "Want", want) 
		 } 
		 } 
		 

func TestNegativeHousePrice(t *testing.T){ 
	 HousePrice := -100000 

	 got := calculateLBTT(HousePrice) 

	 want := 0 

	 if got != want { 
		t.Error("Got", got, "Want", want) 
	}
	 }