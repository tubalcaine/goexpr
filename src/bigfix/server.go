package bigfix

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	Server *BFServer
	Conn   *http.Client
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
func NewBFSession(srv BFServer) (*BFSession, error) {
	var client BFSession

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cli := &http.Client{Transport: customTransport}

	client.Conn = cli
	client.Server = &srv

	return &client, nil
}

// Query - A standalone Session Relevance query using a BFServer. Not
// a persisten session. Useful for rare, one-off queries. Use a session based
// function when you are doing a lot of operations.
func Query(srv BFServer, query *BFSrQuery) (string, error) {
	var apiurl string

	apiurl, _ = BaseBFURL(srv)
	apiurl = apiurl + "/api/query"

	resp, err := http.PostForm(apiurl, url.Values{"relevance": {query.SessionRelevance}})

	if err == nil {
		defer resp.Body.Close()

		result, err := ioutil.ReadAll(resp.Body)

		return string(result), err
	}

	return "", err
}

// SessionQuery - A session based Session Relevance query using a BFServer.
// This uses a persisten session. Useful for many queries in a single program
// or operation.
func SessionQuery(sess *BFSession, query *BFSrQuery) (string, error) {
	var apiurl string

	apiurl, _ = BaseBFURL(*sess.Server)
	apiurl = apiurl + "/api/query"

	data := url.Values{}
	data.Set("relevance", query.SessionRelevance)

	req, err := http.NewRequest("POST", apiurl, strings.NewReader(data.Encode()))

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(sess.Server.Username, sess.Server.Password)

	//	req.Body = url.Values{"relevance": {query.SessionRelevance}}

	resp, err := sess.Conn.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	rtext, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(rtext), nil
}

// MakeSrQuery - Make a query struct from a string
func MakeSrQuery(qstr string) (*BFSrQuery, error) {
	var srq BFSrQuery

	srq.SessionRelevance = qstr
	srq.encodedSessionRelevance = url.QueryEscape(qstr)

	return &srq, nil
}
