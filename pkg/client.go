package pkg

import (
	"net/rpc"
)

type Client struct {
	c *rpc.Client
}

func DialServer() (*Client, error) {
	if c, err := rpc.Dial("tcp", Addr); err != nil {
		return nil, err
	} else {
		return &Client{c: c}, nil
	}
}

func (c *Client) Request(rsp *string) error {
	if err := c.c.Call(ServerName+".Server", "hello, server!", rsp); err != nil {
		return err
	}

	return nil
}
