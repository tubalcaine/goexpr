package bigfix

import (
	"net/http"
)

// BFServer - This structure represents the definition of a BigFix server
// and authentication to same
type BFServer struct {
	Name     string
	Port     string
	Username string
	Password string
}

// BFSession - Represents a persistent connection to a BigFix server
type BFSession struct {
	Conn http.Client
}

// BFSrQuery - Represents a session relevance query
type BFSrQuery struct {
	SessionRelevance        string
	encodedSessionRelevance string
}

// BaseBFURL - returns a URL for the BFServer passed in
func BaseBFURL(srv BFServer) (string, int) {
	return "https://" + srv.Name + ":" + srv.Port, 0
}

// NewBFSession - Create and return a BFSession for a given BFServer
func NewBFSession(srv BFServer) (BFSession, int) {

}

// SessionQuery - A standalone Session Relevance query using a BFServer. Not
// a persisten session. Useful for rare, one-off queries. Use a session based
// function when you are doing a lot of operations.
func SessionQuery(srv BFServer, query BFSrQuery) (string, int) {
	var url string

	url, _ = BaseBFURL(srv) + "/api/query"

	resp, err := http.PostForm()

}

// SessionQuery - A session based Session Relevance query using a BFServer.
// This uses a persisten session. Useful for many queries in a single program
// or operation.
func SessionQuery(sess BFSession, query BFSrQuery) (string, int) {

}
