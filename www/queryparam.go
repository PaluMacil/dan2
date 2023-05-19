package www

import (
	"net/url"
)

func QueryParam(url *url.URL, paramName, defaultValue string) string {
	value := url.Query().Get(paramName)
	if value == "" {
		return defaultValue
	}
	return value
}
