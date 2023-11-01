package tests

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGoENVVariableExistence(t *testing.T) {
	result, exists := os.LookupEnv("GOPROXY")
	if exists {
		assert.Equal(t, "https://proxy.golang.org,direct", result)
	}
}
