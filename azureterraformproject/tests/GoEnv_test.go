package tests

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGoENVVariableExistence(t *testing.T) {
	//log file
	file, err := os.OpenFile("logs/testlog", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	result, exists := os.LookupEnv("GOPROXY")
	if exists {
		assert.Equal(t, "https://proxy.golang.org,direct", result)
	}
	log.SetOutput(file)
	log.Println("Pushing Data to Log File..")
	//read custom environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found...")
	}
	result, exists = os.LookupEnv("cloudName")
	if exists {
		assert.Equal(t, "AzureCloud", result)
	}

}
