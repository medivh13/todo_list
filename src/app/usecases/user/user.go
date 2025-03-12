package user

import (
	"log"

	dto "todo_list/src/app/dto/user"

	repo "todo_list/src/app/repositories/user"
)

type UserUCInterface interface {
	RegisterUser(data *dto.RegisterUserReqDTO) (*dto.RegisterUserRespDTO, error)
	SignIn(data *dto.SignInReqDTO) (*dto.RegisterUserRespDTO, error)
	RefreshToken(data *dto.RefreshTokenReq) (*dto.RefreshTokenResp, error)
}

type UserUseCase struct {
	Repo repo.UserRepository
}

func NewUserUseCase(userRepo repo.UserRepository) *UserUseCase {
	return &UserUseCase{
		Repo: userRepo,
	}
}

func (uc *UserUseCase) RegisterUser(data *dto.RegisterUserReqDTO) (*dto.RegisterUserRespDTO, error) {

	resp, err := uc.Repo.RegisterUser(data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

func (uc *UserUseCase) SignIn(data *dto.SignInReqDTO) (*dto.RegisterUserRespDTO, error) {

	resp, err := uc.Repo.SignIn(data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

func (uc *UserUseCase) RefreshToken(data *dto.RefreshTokenReq) (*dto.RefreshTokenResp, error) {
	resp, err := uc.Repo.RefreshToken(data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
