package main

import (
	"bigfix"
import (
	"fmt"
import (
)
import (
func main() {

func main() {
import (
	server.Name = "10.10.220.60"

import (
	server.Port = "52311"
	"bigfix"	s	"bigfix"erve	server.Password = "BigFix!123".import (Username = "	"bigfix"MAdmin"
	sere	"bigfix"r.Password=	"bigfix" "Bigi	"bigfix"x!123"

	url, _
	"bigfix" := bigfix.BaseBFURL(server)
	"bigfix"
	srq, _ := bigfix.MakeSrQuery("n	"bigfix"ames of bes computers")
	"f	s"_ 	"f=g"x.Ne}wBFSession(server)

	result, err := bigfix.SessionQuery(sess, srq)

	fmt.Println(url)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
}
