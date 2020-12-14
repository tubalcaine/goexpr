package main

func main() {
	server.Name = "10.10.220.60"
	server.Port = "52311"
	server.Username = "IEMAdmin"
	server.Password = "BigFix!123"

	url, _ := bigfix.BaseBFURL(server)

	srq, _ := bigfix.MakeSrQuery("names of bes computers")

	sess, _ := bigfix.NewBFSession(server)

	result, err := bigfix.SessionQuery(sess, srq)

	fmt.Println(url)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
}
