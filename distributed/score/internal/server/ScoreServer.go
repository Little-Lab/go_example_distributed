// Code generated by goctl. DO NOT EDIT.
// Source: scores.proto

package server

import (
	"context"

	"score/internal/logic"
	"score/internal/svc"
	"score/pb/scores"
)

type ScoreServer struct {
	svcCtx *svc.ServiceContext
	scores.UnimplementedScoreServer
}

func NewScoreServer(svcCtx *svc.ServiceContext) *ScoreServer {
	return &ScoreServer{
		svcCtx: svcCtx,
	}
}

// -----------------------score-----------------------
func (s *ScoreServer) AddScore(ctx context.Context, in *scores.AddScoreReq) (*scores.AddScoreResp, error) {
	l := logic.NewAddScoreLogic(ctx, s.svcCtx)
	return l.AddScore(in)
}

func (s *ScoreServer) TranIncScore(ctx context.Context, in *scores.AddScoreReq) (*scores.AddScoreResp, error) {
	l := logic.NewTranIncScoreLogic(ctx, s.svcCtx)
	return l.TranIncScore(in)
}

func (s *ScoreServer) TranDecScore(ctx context.Context, in *scores.AddScoreReq) (*scores.AddScoreResp, error) {
	l := logic.NewTranDecScoreLogic(ctx, s.svcCtx)
	return l.TranDecScore(in)
}

func (s *ScoreServer) UpdateScore(ctx context.Context, in *scores.UpdateScoreReq) (*scores.UpdateScoreResp, error) {
	l := logic.NewUpdateScoreLogic(ctx, s.svcCtx)
	return l.UpdateScore(in)
}

func (s *ScoreServer) DelScore(ctx context.Context, in *scores.DelScoreReq) (*scores.DelScoreResp, error) {
	l := logic.NewDelScoreLogic(ctx, s.svcCtx)
	return l.DelScore(in)
}

func (s *ScoreServer) GetScoreById(ctx context.Context, in *scores.GetScoreByIdReq) (*scores.GetScoreByIdResp, error) {
	l := logic.NewGetScoreByIdLogic(ctx, s.svcCtx)
	return l.GetScoreById(in)
}

func (s *ScoreServer) SearchScore(ctx context.Context, in *scores.SearchScoreReq) (*scores.SearchScoreResp, error) {
	l := logic.NewSearchScoreLogic(ctx, s.svcCtx)
	return l.SearchScore(in)
}
