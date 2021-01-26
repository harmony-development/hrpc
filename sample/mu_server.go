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

func (m *MuServ) MuMute(ctx context.Context, in chan *jiti.Ping, out chan *jiti.Pong, headers http.Header) {
	for {
		data, ok := <-in
		if !ok {
			break
		}

		out <- &jiti.Pong{
			Mu: data.Mu,
		}
	}
}
