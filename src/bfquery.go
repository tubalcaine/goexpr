package main

import (
	"bigfix"
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	var server bigfix.BFServer

	bfqueryApp := cli.NewApp()

	bfqueryApp.Name = "bfquery"
	bfqueryApp.Usage = "Run a session relevance query and write the resulting XML to stdout"
	bfqueryApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Value: "names of bes computers",
			Usage: "Query to run",
		},
		cli.StringFlag{
			Name:  "server, s",
			Value: "10.10.220.59",
			Usage: "Name or address of BigFix root server",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "52311",
			Usage: "TCP port of BigFix root server",
		},
		cli.StringFlag{
			Name:  "user, u",
			Value: "IEMAdmin",
			Usage: "BigFix console user name (for REST API)",
		},
		cli.StringFlag{
			Name:  "password, w",
			Value: "BigFix!123",
			Usage: "BigFix console user password (for REST API)",
		},
	}

	bfqueryApp.Action = func(c *cli.Context) error {
		server.Name = c.GlobalString("server")
		server.Port = c.GlobalString("port")
		server.Username = c.GlobalString("user")
		server.Password = c.GlobalString("password")

		url, _ := bigfix.BaseBFURL(server)

		srq, _ := bigfix.MakeSrQuery(c.GlobalString("query"))

		sess, _ := bigfix.NewBFSession(server)

		result, err := bigfix.SessionQuery(sess, srq)

		fmt.Println(url)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	}

	bfqueryApp.Run(os.Args)
}
