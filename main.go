package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: time.Duration(4 * time.Second),
	}
}

func main() {
	weatherToCache()
	go weatherSchedule()
	iris.Get("/weather", func(c *iris.Context) {
		c.Text(200, GetItemFromCache(weatherURL))
	})
	iris.Listen(":8080")
}

func weatherSchedule() {
	for range time.Tick(time.Hour * 1) {
		weatherToCache()
	}
}

func weatherToCache() {
	fmt.Println("fetching weather")
	if w := fetchWeather(); w != nil {
		j, err := json.Marshal(&w.Query.Results.Channel.Item)
		if err != nil {
			fmt.Println(err)
		} else {
			AddItemToCache(weatherURL, string(j))
		}
	}
}

func fetchWeather() *Weather {
	var w Weather
	resp, err := client.Get(weatherURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err = json.Unmarshal(buf, &w); err != nil {
		fmt.Println(err)
	}
	return &w
}
