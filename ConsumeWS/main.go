package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Error struct {
	Error struct {
		Code    int
		Message string
	}
}

type Result struct {
	Location struct {
		Name    string
		Region  string
		Country string
	}
	Current struct {
		Last_updated string
		Temp         float64
		Condition    struct {
			Text string
		}
		Wind_kph float64
	}
}

var apis map[int]string

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
					// fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					re[API] = result["error"]
					// fmt.Println(result["error"])
				}
				// store the result in channel
				c <- re
				fmt.Println("Result for API 1 stored")
			case 2:
				if result["location"].(map[string]interface{})["name"] != "" {
					re[API] = result
					// fmt.Printf("%+v\n", result)
				} else {
					re[API] = result
					// fmt.Printf("%+v\n", result)
				}
				// store result in channel
				c <- re
				fmt.Println("Result for API 2 stored")
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func usingStruct() {
	url := "http://api.weatherapi.com/v1/current.json?key=97b7ae49dc5c4b08add122134230405&q=Ottawa"
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal([]byte(body), &result)
			if result.Location.Name != "" {
				// fmt.Println(string(body))
				var result Result
				json.Unmarshal([]byte(body), &result)
				fmt.Println(result.Location.Name)
				fmt.Printf("%+v\n", result)
			} else {
				var error Error
				json.Unmarshal([]byte(body), &error)
				fmt.Println(error.Error.Message)
				// fmt.Println(string(body))
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

// channel to store map[int]interface{}
var c chan map[int]interface{}

func main() {
	c = make(chan map[int]interface{})
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=" + "48tlFTkA0SVqYMLHadpzN6MGzFOMFLHP"
	apis[2] = "http://api.weatherapi.com/v1/current.json?key=97b7ae49dc5c4b08add122134230405&q=Ottawa"
	go fetchData(1)
	go fetchData(2)
	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done")
	fmt.Scanln()
}
