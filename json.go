package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type People struct {
	Firstname string
	Lastname  string
	Details   struct {
		Height int
		Weight float32
	}
}

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

type Name struct {
	FirstName string
	LastName  string
}

type Address struct {
	Line1 string
	Line2 string
	Line3 string
}

type Customer struct {
	Name    Name
	Email   string
	Address Address
	DOB     time.Time
}

func decodeJson() {
	var person People

	jsonString := `{"firstname":"Wei-meng","lastname":"Lee"}`
	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	} else {
		fmt.Println(err)
	}

	var persons []People
	jsonString =
		`[
			{
				"firstname":"Anakin",
				"lastname":"Skywalker",
				"details": {
					"height":188,
					"weight":85
				}
			},
			{
				"firstname":"Obiwan",
				"lastname":"Kenobi",
				"details": {
					"height":180,
					"weight":80
				}
			}
		]`

	json.Unmarshal([]byte(jsonString), &persons)

	for _, person := range persons {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
		fmt.Println(person.Details.Height)
		fmt.Println(person.Details.Weight)
	}

	jsonString =
		`{
			"base currency":"EUR",
			"destination currency":"USD"
		}`
	var rates Rates
	json.Unmarshal([]byte(jsonString), &rates)
	fmt.Println(rates.Base)
	fmt.Println(rates.Symbol)

}

func mapJson() {
	jsonString :=
		`{
		"success": true,
		"timestamp": 1588779306,
		"base": "EUR",
		"date": "2020-05-06",
		"rates": {
			"AUD": 1.683349,
			"CAD": 1.528643,
			"GBP": 0.874757,
			"SGD": 1.534513,
			"USD": 1.080054
		}
	}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString), &result)
	fmt.Println(result["success"])

	rates := result["rates"]
	fmt.Println(rates)
	currencies := rates.(map[string]interface{})
	SGD := currencies["SGD"]
	fmt.Println(SGD)
}

func main() {
	decodeJson()
	mapJson()

	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")

	john := Customer{
		Name:  Name{FirstName: "John", LastName: "Smith"},
		Email: "johnsmith@example.com",
		Address: Address{
			Line1: "The White House",
			Line2: "1600 Pennsylvania Avenue NW",
			Line3: "Washington, DC 20500",
		},
		DOB: dob,
	}

	johnJSON, err := json.MarshalIndent(john, "", "	")
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}
