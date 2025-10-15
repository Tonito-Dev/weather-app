package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	location Location `json:"location"`
	current  Current  `json:"current"`
}

type Location struct {
	name    string `json:"name"`
	country string `json:"country"`
}

type Current struct {
	temp_c    int       `json:"temp_c"`
	condition Condition `json:"condition"`
	wind_kph  int       `json:"wind_kph"`
	wind_dir  string    `json:"wind_dir"`
	humidity  int       `json:"humidity"`
}

type Condition struct {
	text string `json:"text"`
}

func main() {
	API_key := "f36279aa4635403f869131542251510"
	response, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + API_key + "&q=London&aqi=no")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
