package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	sites := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://soundcloud.com/",
		"https://www.sofascore.com/",
		"https://qabilah.com/",
		"https://linkedin.com/",
	}
	println("--------- START --------")
	channel := make(chan string)
	for _, value := range sites {
		go getCallback(value, channel)
	}
	for l := range channel {
		go func(link string) {
			time.Sleep(time.Second)
			getCallback(link, channel)
		}(l)
	}
}

func getCallback(value string, channel chan string) {
	_, err := http.Get(value)
	if err != nil {
		println("RESPONSE ERROR: ", err)
		os.Exit(1)
	}
	println(value, " => is up")

	channel <- value
}
