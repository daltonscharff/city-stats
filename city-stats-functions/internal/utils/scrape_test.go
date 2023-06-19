package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapeValidUrl(t *testing.T) {
	body, err := Scrape("https://jsonplaceholder.typicode.com/todos/1")

	assert.Nil(t, err)
	assert.Greater(t, len(body), 0)
}

func TestScrapeInvalidUrl(t *testing.T) {
	_, err := Scrape("https://jsonplaceholder.typicode.com/todos/1000")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "404")
}
