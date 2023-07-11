package rpcfox

import (
	"context"
	"github.com/civet148/jsonrpc"
	"github.com/civet148/log"
	"net/http"
)

type ClientOption struct {
	Url    string      //url to connect e.g ws://127.0.0.1:8080/rpc/v1
	Header http.Header //request header for websocket or http
}

type Client struct {
	opt *ClientOption
	rpc *jsonrpc.MergeClient
}

func NewClient(opt *ClientOption, handlers ...interface{}) *Client {
	if opt == nil {
		log.Panic("client option is empty")
	}
	if opt.Url == "" {
		log.Panic("remote url is empty")
	}
	ctx := context.Background()
	cli, err := jsonrpc.NewMergeClient(ctx, opt.Url, NameSpaceRpcFox, opt.Header, handlers...)
	if err != nil {
		log.Panic("new client error [%s]", err.Error())
	}
	return &Client{
		opt: opt,
		rpc: cli,
	}
}

func (c *Client) Close() {
	c.rpc.Close()
}
