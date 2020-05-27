package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	url string
}

func (c *Client) request() (*map[string]interface{}, error) {
	var m = make(map[string]interface{})

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cl := &http.Client{Transport: tr}
	resp, err := cl.Get(c.url)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("failed to get values " + c.url)
	}
	return &m, json.NewDecoder(resp.Body).Decode(&m)
}

func (c *Client) Get() (*map[string]interface{}, error) {
	var err error
	for i := 0; i < 10; i++ {
		if m, e := c.request(); e == nil {
			return m, nil
		} else {
			err = e
			time.Sleep(1 * time.Second)
		}
	}
	return nil, err
}

func NewClient(client string) *Client {
	return &Client{url: client}
}
