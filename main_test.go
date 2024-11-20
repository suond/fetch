package main

import (
	"encoding/json"
	"testing"
)

func TestCalcTotalPointsWithMorningReceipt(t *testing.T) {
	data := []byte(`
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
`)

	var parsedData Receipt

	json.Unmarshal(data, &parsedData)
	pts := calcTotalPoints(parsedData)
	expected := 15

	if pts != expected {
		t.Errorf("got %q, wanted %q", pts, expected)
	}
}

func TestCalcTotalPointsWithSimpleReceipt(t *testing.T) {
	data := []byte(`
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
`)

	var parsedData Receipt

	json.Unmarshal(data, &parsedData)
	pts := calcTotalPoints(parsedData)
	expected := 31

	if pts != expected {
		t.Errorf("got %q, wanted %q", pts, expected)
	}
}

func TestCalcTotalPointsWithTargetReceipt(t *testing.T) {
	data := []byte(`
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
`)

	var parsedData Receipt

	json.Unmarshal(data, &parsedData)
	pts := calcTotalPoints(parsedData)
	expected := 28

	if pts != expected {
		t.Errorf("got %q, wanted %q", pts, expected)
	}
}

func TestCalcTotalPointsWithMMReceipt(t *testing.T) {
	data := []byte(`
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
`)

	var parsedData Receipt

	json.Unmarshal(data, &parsedData)
	pts := calcTotalPoints(parsedData)
	expected := 109

	if pts != expected {
		t.Errorf("got %q, wanted %q", pts, expected)
	}
}

func TestCalcTotalPointsWithMMModifiedReceipt(t *testing.T) {
	data := []byte(`
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade9",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
`)

	var parsedData Receipt

	json.Unmarshal(data, &parsedData)
	pts := calcTotalPoints(parsedData)
	expected := 105

	if pts != expected {
		t.Errorf("got %q, wanted %q", pts, expected)
	}
}
