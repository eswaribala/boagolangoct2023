package main

import (
	"github.com/hashicorp/vault/api"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {

	vaultAddr := "http://localhost:8200"
	rootToken := "s.Lo4buAvItmJaxiSYeLSyfgcA"
	log.Println("Entered")
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(rootToken)

	//read the credentials from the client
	secrets, error := client.Logical().Read("secret/mysqlsecret")
	if error != nil {
		log.Println("Secret not found....", error)
	}
	for key, value := range secrets.Data {
		log.Println(key, value)
	}

}
