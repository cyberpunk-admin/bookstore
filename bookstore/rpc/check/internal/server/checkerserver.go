// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package server

import (
	"context"

	"check/internal/logic"
	"check/internal/svc"
	"check/pb/check"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
	check.UnimplementedCheckerServer
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *check.CheckReq) (*check.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
