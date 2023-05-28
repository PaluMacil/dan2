package www

import (
	"net/url"
)

// QueryParam gets a query parameter from the url or returns the default value if it can't be found
func QueryParam(url *url.URL, paramName, defaultValue string) string {
	value := url.Query().Get(paramName)
	if value == "" {
		return defaultValue
	}
	return value
}
