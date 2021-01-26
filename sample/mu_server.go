package main

import (
	"context"
	"hrpc/sample/gen/hrpc/jiti"
	"net/http"
)

type MuServ struct {
}

func (m *MuServ) Mu(ctx context.Context, r *jiti.Ping, headers http.Header) (resp *jiti.Pong, err error) {
	return &jiti.Pong{
		Mu: r.Mu,
	}, nil
}
