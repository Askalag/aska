package conv

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

func SignInRequestToUserV1(req *av1.SignInRequest) *repository.User {
	return &repository.User{
		Login:    req.Login,
		Password: req.Password,
	}
}

func SignUpRequestToUserV1(req *av1.SignUpRequest) *repository.User {
	return &repository.User{
		Id:        0,
		Login:     req.Login,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Email:     req.Email,
	}
}
