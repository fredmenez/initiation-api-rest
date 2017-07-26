package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Pattern string
	Handler http.HandlerFunc
}
