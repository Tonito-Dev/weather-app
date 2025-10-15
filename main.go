package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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
