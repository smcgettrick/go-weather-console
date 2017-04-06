package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const baseWeatherURL = "http://api.wunderground.com/api/"

var configuration Configuration

func init() {
	configuration = LoadConfiguration()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter zip code: ")
	zipcode, _ := reader.ReadString('\n')
	zipcode = strings.Replace(zipcode, "\r\n", "", -1)

	if validate(zipcode) {
		weather, err := getWeather(zipcode)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\nLocation: %s, %s\n", weather.Location.City, weather.Location.State)
		fmt.Printf("Current Temperature: %s\n", weather.CurrentObservation.TemperatureString)
		fmt.Printf("Current Observation: %s\n", weather.CurrentObservation.Weather)
	} else {
		fmt.Print("Invalid Zip Code!\n")
		os.Exit(-1)
	}
}

func validate(zipcode string) bool {
	var validZipCode = regexp.MustCompile(`^[0-9]{5}(?:-[0-9]{4})?$`)

	return validZipCode.MatchString(zipcode)
}

func getWeather(zipcode string) (*WeatherData, error) {
	url := baseWeatherURL + configuration.AppID + "/geolookup/conditions/q/PA/" + zipcode + ".json"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var weatherData = new(WeatherData)

	err = json.Unmarshal(body, &weatherData)

	return weatherData, err
}
