package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "World!")
	}
}
