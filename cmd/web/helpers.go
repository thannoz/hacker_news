package main

import (
	"net/http"
	"strconv"
)

func (a *application) readInt(r *http.Request, key string) int {
	v, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return 0
	}
	return v
}
