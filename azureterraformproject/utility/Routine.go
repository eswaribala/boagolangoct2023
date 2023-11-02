package main

import (
	"fmt"
	"log"
	"net/http"
)

func routineAccessLink(link *string, pChannel chan string) {
	response, err := http.Get(*link)
	if err != nil {
		log.Fatal("Site Not Available")
	} else {
		log.Printf("Found the site=%s and having response code%d\n",
			*link, response.StatusCode)
		pChannel <- "Routine completed the task"
	}

}

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.portal.azure.com",
		"https://www.rediffmail.com",
	}

	//create the channel
	channel := make(chan string)

	for _, value := range links {
		//function call converted to routine call
		go routineAccessLink(&value, channel)
	}

	for _, link := range links {
		fmt.Printf("Message received from the link=%s to channel%s\n", link, <-channel)
	}

	log.Println("Visited all the links")

}
