package v1

import "context"
import "net/http"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/harmony-development/legato/gen/auth/v1"

type AuthServiceServer interface {
	Federate(ctx context.Context, r *v1.FederateRequest) (resp v1.FederateRequest, err error)

	LoginFederated(ctx context.Context, r *v1.LoginFederatedRequest) (resp v1.LoginFederatedRequest, err error)

	Key(ctx context.Context, r *empty.Empty) (resp v1.Empty, err error)

	BeginAuth(ctx context.Context, r *empty.Empty) (resp v1.Empty, err error)

	NextStep(ctx context.Context, r *v1.NextStepRequest) (resp v1.NextStepRequest, err error)

	StepBack(ctx context.Context, r *v1.StepBackRequest) (resp v1.StepBackRequest, err error)
}

type AuthServiceHandler struct {
	Server AuthServiceServer
}

func (h *AuthServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.auth.v1.AuthService/Federate":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/LoginFederated":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/Key":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/BeginAuth":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/NextStep":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/StepBack":
		{
			panic("unimplemented")
		}

	case "/protocol.auth.v1.AuthService/StreamSteps":
		{
			panic("unimplemented")
		}

	}
}
