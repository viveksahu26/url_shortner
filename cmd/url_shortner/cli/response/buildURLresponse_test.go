package response_test

import (
	"testing"

	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli/response"
	"gotest.tools/assert"
)

func TestBuildResonse(t *testing.T) {
	expected := response.BuildURLWithResponse("localhost:8080", "duj3P^+d*C", "http://viveksahu.com")
	actual := response.URLResponse{
		OriginalURL: "http://viveksahu.com",
		ShortURL:    "http://localhost:8080/duj3P^+d*C",
	}
	assert.Equal(t, expected.ShortURL, actual.ShortURL)
}
