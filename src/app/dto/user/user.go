package user

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterUserReqDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *RegisterUserReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password,
			validation.Required,
			validation.Length(8, 15),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>_\[\]\\\/]`)).Error("must contain at least one special character"),
		),
		validation.Field(&dto.Name, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterUserRespDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *SignInReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type SignInModelDTO struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	Token string `json:"token"`
}
