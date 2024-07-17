package main

import (
	"errors"
	"testing"
)

func TestLBTTBelowFirstBand(t *testing.T) {
	housePrice := 100000

	got, _ := calculateLBTT(housePrice)

	want := 0

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTAtFirstBandLimit(t *testing.T) {
	housePrice := 145000

	got, _ := calculateLBTT(housePrice)

	want := 0

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTInSecondBand(t *testing.T) {
	housePrice := 180000

	got, _ := calculateLBTT(housePrice)

	want := 700

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTAtSecondBandLimit(t *testing.T) {
	housePrice := 250000
	
	got, _ := calculateLBTT(housePrice)

	want := 2100

	if got != want {

		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTInThirdBand(t *testing.T) {
	housePrice := 300000

	got, _ := calculateLBTT(housePrice)

	want := 4600

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTAtThirdBandLimit(t *testing.T) {
	housePrice := 325000

	got, _ := calculateLBTT(housePrice)

	want := 5850

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTInFourthBand(t *testing.T) {
	housePrice := 600000

	got, _ := calculateLBTT(housePrice)

	want := 33350

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTAtFourthBandLimit(t *testing.T) {
	housePrice := 750000

	got, _ := calculateLBTT(housePrice)

	want := 48350

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTAboveFourthBand(t *testing.T) {
	housePrice := 900000

	got, _ := calculateLBTT(housePrice)

	want := 66350

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestLBTTSpecificSpotPrice(t *testing.T) {
	housePrice := 145550

	got, _ := calculateLBTT(housePrice)

	want := 11

	if got != want {
		t.Error("Got", got, "Want", want)
	}
}

func TestTwoPercentCalculation(t *testing.T){
	HousePrice:=550

	got:=TwoPercentCalculation(HousePrice)

	want:=11

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
	HousePrice:=100

	got:=TenPercentCalculation(HousePrice)

	want:=10

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestTwelvePercentCalculation(t *testing.T){
	HousePrice:=100

	got:=TwelvePercentCalculation(HousePrice)

	want:=12

	if got!=want{
		t.Error("Got",got,"Want",want,)
	}
}

func TestNegativeHousePrice(t *testing.T){ 
	HousePrice := -1000000 

	_, gotErr:= calculateLBTT(HousePrice) 

	wantErr := errors.New("cannot be negative houseprice") 

	if gotErr.Error() != wantErr.Error() { 
	   t.Error("Got", gotErr, "Want", wantErr) 
   }

}