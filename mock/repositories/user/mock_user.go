package user

import (
	dto "todo_list/src/app/dto/user"
	repo "todo_list/src/app/repositories/user"

	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func NewMockUser() *MockUser {
	return &MockUser{}
}

var _ repo.UserRepository = &MockUser{}

func (o *MockUser) RegisterUser(data *dto.RegisterUserReqDTO) (*dto.RegisterUserRespDTO, error) {
	args := o.Called(data)

	var (
		resp *dto.RegisterUserRespDTO
		err  error
	)

	if n, ok := args.Get(0).(*dto.RegisterUserRespDTO); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}

func (o *MockUser) SignIn(data *dto.SignInReqDTO) (*dto.RegisterUserRespDTO, error) {
	args := o.Called(data)

	var (
		resp *dto.RegisterUserRespDTO
		err  error
	)

	if n, ok := args.Get(0).(*dto.RegisterUserRespDTO); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}

func (o *MockUser) RefreshToken(data *dto.RefreshTokenReq) (*dto.RefreshTokenResp, error) {
	args := o.Called(data)

	var (
		resp *dto.RefreshTokenResp
		err  error
	)

	if n, ok := args.Get(0).(*dto.RefreshTokenResp); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}
