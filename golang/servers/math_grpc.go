package servers

import (
	"context"

	"github.com/marlonklc/poc-grpc/gopb"
)

type Math struct {
}

func (m *Math) Sum(ctx context.Context, req *gopb.SumRequest) (*gopb.SumResponse, error) {

	res := req.Sum.A + req.Sum.B

	return &gopb.SumResponse{
		Result: res,
	}, nil
}
