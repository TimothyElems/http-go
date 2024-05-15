/*
Test with Insomnia
*/

package main

import (
	"net/http"

	"github.com/TimothyElems/http-go.git/api"
)

func main () {
	srv := api.NewServer()
	http.ListenAndServe(":8000", srv)
}