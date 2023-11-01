package tests

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGoENVVariableExistence(t *testing.T) {
	result, exists := os.LookupEnv("GOPROXY")
	if exists {
		assert.Equal(t, "https://proxy.golang.org,direct", result)
	}

	//log file
	file, err := os.OpenFile("testlog", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(file)
	log.Println("Pushing Data to Log File..")

}
