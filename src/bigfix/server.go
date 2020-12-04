package bigfix

// This structure represents the definition of a BigFix server
// and authentication to same
type BFServer struct {
	Name     string
	Port     string
	Username string
	Password string
}

// This function returns the base URL of the server structure passed in
func BaseBFURL(srv BFServer) (string, int) {
	return "https://" + srv.Name + ":" + srv.Port, 0
}
