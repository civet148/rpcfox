package main

import (
	"context"
	"github.com/civet148/log"
	"github.com/civet148/rpcfox.git"
	"github.com/civet148/rpcfox.git/examples/proto"
)

type MessageClient struct {
	UserLogin     func(ctx context.Context, req *proto.UserLoginReq) (resp *proto.UserLoginResp, err error)   //用户登录
	UserLogout    func(ctx context.Context, req *proto.UserLogoutReq) (resp *proto.UserLogoutResp, err error) //用户退出
	SendNotify    func(ctx context.Context, req *proto.SendNotifyReq) (err error)                             //只发送一个通知（服务端无需返回任何数据）
	UserSubscribe func(ctx context.Context) (notifies <-chan []*proto.UserNotify, err error)                  //发起订阅,监听用户事件
}

func main() {
	var mc MessageClient
	c := rpcfox.NewClient(&rpcfox.ClientOption{
		Url:    proto.ServerUrl,
		Header: nil,
	}, &mc)

	defer c.Close()

	ctx := context.TODO()
	var err error
	var login *proto.UserLoginResp
	login, err = mc.UserLogin(ctx, &proto.UserLoginReq{
		UserName: "lory",
		Password: "123456",
	})
	if err != nil {
		log.Errorf("UserLogin return error [%s]", err.Error())
		return
	}
	log.Infof("UserLogin return [%+v]", login)

	var logout *proto.UserLogoutResp
	logout, err = mc.UserLogout(ctx, &proto.UserLogoutReq{
		UserId: login.UserId,
		Token:  login.Token,
	})
	if err != nil {
		log.Errorf("UserLogout return error [%s]", err.Error())
		return
	}
	log.Infof("UserLogout return [%+v]", logout)

	err = mc.SendNotify(ctx, &proto.SendNotifyReq{
		UserId:  login.UserId,
		Message: "Hello ~",
	})
	if err != nil {
		log.Errorf("SendNotify return error [%s]", err.Error())
		return
	}
	log.Infof("SendNotify ok")

	notifies, err := mc.UserSubscribe(ctx)
	if err != nil {
		log.Errorf("UserSubscribe return error [%s]", err.Error())
		return
	}
	for {
		select {
		case notis := <-notifies:
			for _, v := range notis {
				log.Infof("subscribe event notify [%+v]", v)
			}
		}
	}
	//exit := make(chan bool, 1)
	//<-exit
}
