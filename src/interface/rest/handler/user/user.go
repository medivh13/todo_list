package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	dto "todo_list/src/app/dto/user"
	usecases "todo_list/src/app/usecases/user"
	common_error "todo_list/src/infra/errors"
	"todo_list/src/infra/helper"
	"todo_list/src/interface/rest/response"

	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
)

type UserHandlerInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	response response.IResponseClient
	usecase  usecases.UserUCInterface
}

func NewUserHandler(r response.IResponseClient, h usecases.UserUCInterface) UserHandlerInterface {
	return &UserHandler{
		response: r,
		usecase:  h,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.RegisterUserReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.RegisterUser(&postDTO)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			h.response.HttpError(w, common_error.NewError(common_error.USER_ALREADY_EXIST, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Register New User",
		data,
		nil,
	)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.SignInReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.SignIn(&postDTO)
	if err != nil {
		log.Println(err)
		if err.Error() == "no rows returned from the query" {
			h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	h.response.JSON(
		w,
		"Successful SignIn User",
		data,
		nil,
	)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	req := dto.RefreshTokenReq{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	// Validasi refresh token
	_, err = helper.VerifyToken(req.RefreshToken)
	if err != nil {

		if err == jwt.ErrSignatureInvalid || strings.Contains(err.Error(), "refresh token is expired") {
			h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("refresh_token_expired")))
			return
		}

		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Generate access token baru
	data, err := h.usecase.RefreshToken(&req)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Refresh Token",
		data,
		nil,
	)
}
