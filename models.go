package main

import (
	"strings"
	"time"
)

// structs and methods

type HomeResponse struct {
	Message string `json:"message"`
}

type ProcessResponse struct {
	Id string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

type Receipt struct {
	//retailer, purchase date, purchaseTime, items, total
	Retailer     string `json:"retailer"`
	PurchaseDate PDate  `json:"purchaseDate"`
	PurchaseTime PTime  `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type PDate time.Time
type PTime time.Time

func (cd *PDate) UnmarshalJSON(data []byte) error {
	str := string(data)
	trimmed := strings.Trim(str, `"`)
	t, err := time.Parse("2006-01-02", trimmed)
	if err != nil {
		return err
	}
	*cd = PDate(t)
	return nil
}

func (cd PDate) String() string {
	return time.Time(cd).Format("2006-01-02")
}

func (cd *PTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	trimmed := strings.Trim(str, `"`)
	t, err := time.Parse(`15:04`, trimmed)
	if err != nil {
		return err
	}
	*cd = PTime(t)
	return nil
}

func (cd PTime) String() string {
	return time.Time(cd).Format("15:04")
}
