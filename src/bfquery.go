package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	var server BFServer

	server.name = "10.10.220.60"

	fmt.Println(URL(server))
}
