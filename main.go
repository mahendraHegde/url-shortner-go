package main

import (
	"github.com/mahendrahegde/url-shortner-golang/api"
)

var server = api.Server{}

func main() {
	server.Start()
}
