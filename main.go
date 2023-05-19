package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type weathermonitoring struct {
	DeviceId       string `json:"deviceid"`
	DeviceName     string `json:"devicename"`
	AirTemparature string `json:"airtemparature"`
	Airpressure    string `json:"airpressure"`
	Humidity       string `json:"humidity"`
	WindSpeed      string `json:"windspeed"`
	Rain           string `json:"rain"`
}

var weathervalues = []weathermonitoring{
	{DeviceId: "1", DeviceName: "device1", AirTemparature: "-194.4", Airpressure: "101,325 Pa", Humidity: "9%", WindSpeed: "80 km/h", Rain: "6"},
	{DeviceId: "2", DeviceName: "device2", AirTemparature: "-184.4", Airpressure: "102,325 Pa", Humidity: "8%", WindSpeed: "70 km/h", Rain: "5"},
	{DeviceId: "3", DeviceName: "device3", AirTemparature: "-174.4", Airpressure: "107,325 Pa", Humidity: "7%", WindSpeed: "60 km/h", Rain: "4"},
}

func main() {
	router := gin.Default()
	router.GET("/weathervalues", weather)
	router.GET("/weathervalues/:id", getWeatherById)
	router.POST("/weathervalues", postWeather)

	router.Run("localhost:8080")
}

func weather(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, weathervalues)
}

func postWeather(c *gin.Context) {
	var newValue weathermonitoring

	if err := c.BindJSON(&newValue); err != nil {
		return
	}

	weathervalues = append(weathervalues, newValue)
	c.IndentedJSON(http.StatusCreated, newValue)
}

func getWeatherById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range weathervalues {
		if a.DeviceId == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "device id not found"})

}
