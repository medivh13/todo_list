package user

import (
	"errors"
	mockRepo "todo_list/mock/repositories/user"

	"testing"
	dto "todo_list/src/app/dto/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserUseCase struct {
	mock.Mock
}

type UserUseCaseList struct {
	suite.Suite

	useCase             UserUCInterface
	mockRepo            *mockRepo.MockUser
	dtoRegister         *dto.RegisterUserReqDTO
	dtoSignIn           *dto.SignInReqDTO
	dtoRefreshToken     *dto.RefreshTokenReq
	dtoRegisterResp     *dto.RegisterUserRespDTO
	dtoRefreshTokenResp *dto.RefreshTokenResp
}

func (suite *UserUseCaseList) SetupTest() {

	suite.mockRepo = new(mockRepo.MockUser)
	suite.useCase = NewUserUseCase(suite.mockRepo)

	suite.dtoRegister = &dto.RegisterUserReqDTO{
		Name:     "backend magang",
		Email:    "backendmagang@gmail.com",
		Password: "Magang123_",
	}

	suite.dtoSignIn = &dto.SignInReqDTO{
		Email:    "backendmagang@gmail.com",
		Password: "Magang123_",
	}

	suite.dtoRefreshToken = &dto.RefreshTokenReq{
		RefreshToken: "pokoknyaTokenJwt",
	}

	suite.dtoRegisterResp = &dto.RegisterUserRespDTO{}

	suite.dtoRefreshTokenResp = &dto.RefreshTokenResp{}

}

func (u *UserUseCaseList) TestRegisterUserSuccess() {

	u.mockRepo.Mock.On("RegisterUser", u.dtoRegister).Return(u.dtoRegisterResp, nil)
	_, err := u.useCase.RegisterUser(u.dtoRegister)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestRegisterUserFail() {

	u.mockRepo.Mock.On("RegisterUser", u.dtoRegister).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.RegisterUser(u.dtoRegister)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestSignInSuccess() {

	u.mockRepo.Mock.On("SignIn", u.dtoSignIn).Return(u.dtoRegisterResp, nil)
	_, err := u.useCase.SignIn(u.dtoSignIn)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestSignInFail() {

	u.mockRepo.Mock.On("SignIn", u.dtoSignIn).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.SignIn(u.dtoSignIn)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestRefreshTokenSuccess() {

	u.mockRepo.Mock.On("RefreshToken", u.dtoRefreshToken).Return(u.dtoRefreshTokenResp, nil)
	_, err := u.useCase.RefreshToken(u.dtoRefreshToken)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestRefreshTokenFail() {

	u.mockRepo.Mock.On("RefreshToken", u.dtoRefreshToken).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.RefreshToken(u.dtoRefreshToken)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UserUseCaseList))
}
