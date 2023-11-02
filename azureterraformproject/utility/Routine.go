package main

import (
	"log"
	"net/http"
)

func routineAccessLink(link *string) {
	response, err := http.Get(*link)
	if err != nil {
		log.Fatal("Site Not Available")
	} else {
		log.Printf("Found the site=%s and having response code%d",
			*link, response.StatusCode)
	}

}

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.portal.azure.com",
		"https://www.rediffmail.com",
	}

	for _, value := range links {
		//function call converted to routine call
		go routineAccessLink(&value)
	}

	log.Println("Visited all the links")

}
