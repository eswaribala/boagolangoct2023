package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func routineAccessLink(link *string, pChannel chan string) {
	response, err := http.Get(*link)
	if err != nil {
		log.Fatal("Site Not Available")
	} else {
		log.Printf("Found the site=%s and having response code%d\n",
			*link, response.StatusCode)
		pChannel <- "Routine completed the task" + *link
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
		//create the delay between routines
		time.Sleep(5 * time.Second)
	}

	fmt.Printf("Message received from  to channel%s\n", <-channel)
	fmt.Printf("Message received from  to channel%s\n", <-channel)
	fmt.Printf("Message received from  to channel%s\n", <-channel)
	//fmt.Printf("Message received from  to channel%s\n", <-channel)

	log.Println("Visited all the links")

}
