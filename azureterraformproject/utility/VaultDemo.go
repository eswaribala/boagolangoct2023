package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {

	vaultAddr := "http://127.0.0.1:8200"
	rootToken := "s.1hcA47eSwj5sHHgrbuJQv80F"
	log.Println("Entered")
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(rootToken)

	//read the credentials from the client
	secrets, error := client.Logical().Read("/secret/data/mysqlsecret")
	if error != nil {
		log.Println("Secret not found....", error)
	}
	for _, value := range secrets.Data {

		for k, v := range value.(map[string]interface{}) {
			if k == "username" || k == "password" {
				fmt.Println(v)
			}

		}
	}

}
