package main

import (
	"testing"
)

func TestLBTTFirstBand(t *testing.T){
	HousePrice:= 100000

	got:=calculateLBTT(HousePrice)

	want:=0

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsFirstBand(t *testing.T){
	HousePrice:=145000

	got:=calculateLBTT(HousePrice)

	want:=0

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}


func TestLBTTSecondBand(t *testing.T){
	HousePrice:= 180000

	got:= calculateLBTT(HousePrice)

	want:=700

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsSecondVBand(t *testing.T){
	HousePrice:= 250000

	got:= calculateLBTT(HousePrice)

	want:=2100

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTThirdBand(t *testing.T){
	HousePrice:=300000
	got:= calculateLBTT(HousePrice)

	want:= 4600

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsThirdBand(t *testing.T){
	HousePrice:=325000
	got:= calculateLBTT(HousePrice)

	want:=5850

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTFourthBand(t *testing.T){
	HousePrice:=600000
	got:= calculateLBTT(HousePrice)

	want:= 33350

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestLBTTIsFourthBand(t *testing.T){
	HousePrice:=750000
	got:= calculateLBTT(HousePrice)

	want:=48350

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

func TestFivePercentCalculation(t *testing.T){
	HousePrice:= 100

	got:= FivePercentCalculation(HousePrice)

	want:=5

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}

}

func TestTenPercentCalculation(t *testing.T){
	HousePrice:= 100

	got:= TenPercentCalculation(HousePrice)

	want:= 10

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