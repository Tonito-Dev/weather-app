package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	Temp_c    float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
	Wind_kph  float64   `json:"wind_kph"`
	Wind_dir  string    `json:"wind_dir"`
	Humidity  int       `json:"humidity"`
}

type Condition struct {
	Text string `json:"text"`
}

func main() {
	API_key := "f36279aa4635403f869131542251510"
	response, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + API_key + "&q=Mbeya&aqi=no")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("Region: ", responseObject.Location.Name)
	fmt.Println("Country: ", responseObject.Location.Country)
	fmt.Println("Temp: ", responseObject.Current.Temp_c, "Celsius")
	fmt.Println("Condition: ", responseObject.Current.Condition.Text)
}
