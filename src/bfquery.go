package main

import (
	"bigfix"
	"fmt"
)

func main() {
	var server bigfix.BFServer

	server.Name = "10.10.220.60"
	server.Port = "52311"
	server.Username = "IEMAdmin"
	server.Password = "BigFix!123"

	url, _ := bigfix.BaseBFURL(server)

	result, err := bigfix.Query(server, "names of bes computers")

	fmt.Println(url)
}
