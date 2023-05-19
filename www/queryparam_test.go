package www_test

import (
	"github.com/PaluMacil/dan2/www"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func testQueryParam(t *testing.T, testURL, expectedValue string) {
	parsedUrl, err := url.Parse(testURL)
	assert.NoError(t, err, "parsing test url: %s", testURL)
	actualValue := www.QueryParam(parsedUrl, "this", "maybe")
	assert.Equal(t, expectedValue, actualValue)
}

func TestQueryParam0(t *testing.T) {
	testQueryParam(t, "https://danwolf.net?this=yes&not_this=no", "yes")
}

func TestQueryParam1(t *testing.T) {
	testQueryParam(t, "https://danwolf.net", "maybe")
}

func TestQueryParam2(t *testing.T) {
	testQueryParam(t, "https://danwolf.net?not_this=no", "maybe")
}

func TestQueryParam3(t *testing.T) {
	testQueryParam(t, "https://danwolf.net?not_this=no&this=yes", "yes")
}
