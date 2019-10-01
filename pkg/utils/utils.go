package utils

import "net/http"

// Get a value corresponding to the given key in a query
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
