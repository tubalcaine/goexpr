package bigfix

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
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
	return "https://" + srv.Username + ":" + srv.Password + "@" + srv.Name + ":" + srv.Port, 0
}

// NewBFSession - Create and return a BFSession for a given BFServer
func NewBFSession(srv BFServer) (*BFSession, error) {
	var client BFSession

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cli := &http.Client{Transport: customTransport}

	client.Conn = cli
	return &client, nil
}

// Query - A standalone Session Relevance query using a BFServer. Not
// a persisten session. Useful for rare, one-off queries. Use a session based
// function when you are doing a lot of operations.
func Query(srv BFServer, query BFSrQuery) (string, error) {
	var apiurl string

	apiurl, _ = BaseBFURL(srv)
	apiurl = apiurl + "/api/query"

	resp, err := http.PostForm(apiurl, url.Values{"relevance": {query.SessionRelevance}})
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)

	return string(result), err
}

// SessionQuery - A session based Session Relevance query using a BFServer.
// This uses a persisten session. Useful for many queries in a single program
// or operation.
func SessionQuery(sess BFSession, query BFSrQuery) (string, error) {
	return "", nil
}
