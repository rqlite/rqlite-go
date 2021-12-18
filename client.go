package rqlite

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Node struct {
	u *url.URL
	c *http.Client

	statusURL string
	queryURL string
	executeURL string
}

func NewNode(u *url.URL) *Node {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Node{
		u: u,
		c: c,
		statusURL: u.String() + "/status",
		queryURL: u.String() + "/db/query",
		executeURL: u.String() + "/db/execute",
	}
}

func (n *Node) Status() (io.ReadCloser, error) {
	resp, err := n.c.Get(n.statusURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status endpoint not found")
	}
	return resp.Body, nil
}