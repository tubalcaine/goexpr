package main

import (
	"bigfix"
	"fmt"
)

func main() {
	var server bigfix.BFServer

	server.Name = "10.10.220.60"
	server.Port = "52311"

	url, _ := bigfix.BaseBFURL(server)

	fmt.Println(url)
}
