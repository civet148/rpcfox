package rpcfox

import (
	"github.com/civet148/jsonrpc"
	"github.com/civet148/log"
)

type ServerOption struct {
	Url            string //url to listen e.g ws://127.0.0.1:8080/rpc/v1
	MaxRequestSize int64  //max request body size, default 100MiB
	PingInterval   int    //server ping interval, default 5 seconds
}

type Server struct {
	opt *ServerOption
	rpc *jsonrpc.MergeServer
}

func NewServer(opt *ServerOption, handler interface{}) *Server {
	if opt == nil {
		log.Panic("client option is empty")
	}
	if opt.Url == "" {
		log.Panic("remote url is empty")
	}
	so := &jsonrpc.Option{
		MaxRequestSize: opt.MaxRequestSize,
		PingInterval:   opt.PingInterval,
	}
	s := jsonrpc.NewMergeServer(NameSpaceRpcFox, handler, so)
	return &Server{
		rpc: s,
		opt: opt,
	}
}

func (s *Server) Listen() error {
	log.Infof("url %s listening...", s.opt.Url)
	if err := s.rpc.ListenAndServe(s.opt.Url); err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	//TODO
}
