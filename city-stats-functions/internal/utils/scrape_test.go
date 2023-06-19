package utils

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestScrapeValidUrl(t *testing.T) {
	defer gock.Off()
	gock.New("https://jsonplaceholder.typicode.com").
		Get("/todos/1").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	body, err := Scrape("https://jsonplaceholder.typicode.com/todos/1")

	assert.Nil(t, err)
	assert.Greater(t, len(body), 0)
}

func TestScrapeInvalidUrl(t *testing.T) {
	defer gock.Off()
	gock.New("https://jsonplaceholder.typicode.com").
		Get("/todos/1000").
		Reply(404)

	_, err := Scrape("https://jsonplaceholder.typicode.com/todos/1000")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "404")
}
