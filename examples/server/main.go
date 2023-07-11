package main

import (
	"context"
	"errors"
	"github.com/civet148/log"
	"github.com/civet148/rpcfox.git"
	"github.com/civet148/rpcfox.git/examples/proto"
	"time"
)

const (
	USER_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2ODg5OTU0NTgsImNsYWltX2lhdCI6MTY4ODk1MjI1OCwiY2xhaW1faWQiOjE2LCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibGliaW52aWV3ZXIifQ.sixC2k8a3IistNRMf8nq4gA7M4QdyeXFEhqoaTwh8rI"
)

type MessageAPI interface {
	UserLogin(context.Context, *proto.UserLoginReq) (*proto.UserLoginResp, error)     //用户登录
	UserLogout(context.Context, *proto.UserLogoutReq) (*proto.UserLogoutResp, error)  //用户退出
	SendNotify(context.Context, *proto.SendNotifyReq) error                           //只发送一个通知（服务端无需返回任何数据）
	UserSubscribe(ctx context.Context) (notifies chan []*proto.UserNotify, err error) //用户发起订阅,监听用户事件
}

type MessageServer struct {
}

func (m *MessageServer) UserLogin(ctx context.Context, req *proto.UserLoginReq) (resp *proto.UserLoginResp, err error) {
	if req.UserName == "" || req.Password == "" {
		return nil, errors.New("user or password invalid")
	}
	// database CRUD ...
	return &proto.UserLoginResp{
		UserId:      10010,
		UserName:    "lory",
		Age:         20,
		HomeAddress: "shenzhen china",
		Token:       USER_TOKEN,
	}, nil
}

func (m *MessageServer) UserLogout(ctx context.Context, req *proto.UserLogoutReq) (resp *proto.UserLogoutResp, err error) {
	if req.UserId == 0 {
		return nil, errors.New("user id is empty")
	}
	if req.Token != USER_TOKEN {
		return nil, errors.New("token invalid")
	}
	// database CRUD ...
	return &proto.UserLogoutResp{
		OK: true,
	}, nil
}

func (m *MessageServer) SendNotify(ctx context.Context, req *proto.SendNotifyReq) (err error) {
	log.Infof("user id [%v] notification [%s]", req.UserId, req.Message)
	return nil
}

func (m *MessageServer) UserSubscribe(ctx context.Context) (notifies chan []*proto.UserNotify, err error) {
	log.Infof("in")
	var events []*proto.UserNotify
	events = append(events, &proto.UserNotify{
		UserId:    102328,
		UserEvent: proto.UserEvent_Login,
	})
	events = append(events, &proto.UserNotify{
		UserId:    102329,
		UserEvent: proto.UserEvent_Logout,
	})
	notifies = make(chan []*proto.UserNotify, 50)
	go func() {
		for {
			notifies <- events
			log.Infof("send notification events (%v)", len(events))
			time.Sleep(5 * time.Second)
		}
	}()
	return
}

func main() {
	var handler MessageServer
	var strUrl = proto.ServerUrl
	s := rpcfox.NewServer(&rpcfox.ServerOption{
		Url:            strUrl,
		MaxRequestSize: 0,
		PingInterval:   0,
	}, &handler)
	defer s.Close()

	err := s.Listen()
	if err != nil {
		log.Errorf("listen url %s error [%s]", strUrl, err.Error())
		return
	}
	exit := make(chan bool, 1)
	<-exit
}
