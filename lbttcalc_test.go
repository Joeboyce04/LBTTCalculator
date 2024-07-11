package main

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

	want:=700

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsSecondValue(t *testing.T){
	HousePrice:= 250000

	got:= calculateLBTT(HousePrice)

	want:=2100

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}


func TestTwoPercentCalculation(t *testing.T){
	HousePrice:=100

	got:=TwoPercentCalculation(HousePrice)

	want:=2

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