package logic

import (
	"context"

	"zerovalue/internal/svc"
	"zerovalue/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZerovalueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZerovalueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZerovalueLogic {
	return &ZerovalueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZerovalueLogic) Zerovalue(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Name = req.Name

	return resp, nil
}
