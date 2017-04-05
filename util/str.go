package util

import (
	"log"
	"net/url"
)

// SetQuery set query param 'key' to 'value', it will overwrite existing one
func SetQuery(baseURL string, key, value string) (fullURL string, err error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("parse %s to url error: %v\n", baseURL, err)
		return
	}

	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()

	fullURL = u.String()
	return
}
