package src

import (
	"testing"

	"gotest.tools/assert"
)

func TestBuildResonse(t *testing.T) {
	expected := BuildURLWithResponse("localhost:8080", "duj3P^+d*C", "http://viveksahu.com")
	actual := &URLResponse{
		OriginalURL: "http://viveksahu.com",
		ShortURL:    "http://localhost:8080/duj3P^+d*C",
	}
	assert.Equal(t, expected.ShortURL, actual.ShortURL)
}
