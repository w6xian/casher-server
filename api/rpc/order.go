package rpc

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type RegisterRequest struct {
	Name     string
	Password string
}

type LoginRequest struct {
	Name     string
	Password string
}

type RegisterReply struct {
	Code      int
	AuthToken string
}

type LoginReply struct {
	Code      int
	AuthToken string
}

func (c *Order) Register(ctx context.Context, req *RegisterRequest, reply *RegisterReply) error {
	md5Sum := MD5(req.Name + req.Password)
	reply.Code = 0
	reply.AuthToken = string(md5Sum[:]) + "server Register"
	return nil
}

func (c *Order) Login(ctx context.Context, req *LoginRequest, reply *LoginReply) error {
	md5Sum := MD5(req.Name + req.Password)
	reply.Code = 0
	reply.AuthToken = string(md5Sum[:]) + "server Login"
	return nil
}

func MD5(input string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
