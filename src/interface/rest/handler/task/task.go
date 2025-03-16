package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	dto "todo_list/src/app/dto/task"
	usecases "todo_list/src/app/usecases/task"
	common_error "todo_list/src/infra/errors"
	"todo_list/src/infra/helper"
	"todo_list/src/interface/rest/response"

	"github.com/golang-jwt/jwt"
)

// TaskHandlerInterface mendefinisikan kontrak untuk handler task
type TaskHandlerInterface interface {
	AddTask(w http.ResponseWriter, r *http.Request)
	FinishTask(w http.ResponseWriter, r *http.Request)
	GetTaskList(w http.ResponseWriter, r *http.Request)
}

// TaskHandler adalah implementasi dari TaskHandlerInterface
type TaskHandler struct {
	response response.IResponseClient // Untuk menangani response HTTP
	usecase  usecases.TaskUCInterface // Menghubungkan ke layer use case
}

// NewTaskHandler membuat instance baru dari TaskHandler
func NewTaskHandler(r response.IResponseClient, h usecases.TaskUCInterface) TaskHandlerInterface {
	return &TaskHandler{
		response: r,
		usecase:  h,
	}
}

// extractBearerToken mengekstrak token dari header Authorization
func (h *TaskHandler) extractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization token")
	}

	// Pastikan format header "Authorization" adalah "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil // Mengembalikan token tanpa kata "Bearer"
}

// AddTask menangani request untuk menambahkan task baru
func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	// Ekstrak token dari header Authorization
	tokenString, err := h.extractBearerToken(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Verifikasi token JWT
	dataClaims, err := helper.VerifyToken(tokenString)
	if err != nil {
		// Tangani error jika token tidak valid atau kadaluarsa
		if err == jwt.ErrSignatureInvalid || strings.Contains(err.Error(), "token is expired") {
			h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("token expired")))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Inisialisasi DTO untuk task baru
	postDTO := dto.CreateTaskReqDTO{
		UserID: dataClaims.UserID, // Ambil UserID dari token yang telah diverifikasi
	}

	// Decode body request ke DTO
	err = json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	// Validasi input data task
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	// Panggil use case untuk menambahkan task
	err = h.usecase.AddTask(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	// Beri response sukses
	h.response.JSON(
		w,
		"task baru sedang di proses",
		nil,
		nil,
	)
}

// FinishTask menangani request untuk menyelesaikan task
func (h *TaskHandler) FinishTask(w http.ResponseWriter, r *http.Request) {
	// Ekstrak token dari header Authorization
	tokenString, err := h.extractBearerToken(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Verifikasi token JWT
	_, err = helper.VerifyToken(tokenString)
	if err != nil {
		// Tangani error jika token tidak valid atau kadaluarsa
		if err == jwt.ErrSignatureInvalid || strings.Contains(err.Error(), "token is expired") {
			h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("token expired")))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Inisialisasi DTO untuk menyelesaikan task
	postDTO := dto.FinishtTaskReqDTO{}

	// Decode body request ke DTO
	err = json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	// Panggil use case untuk menyelesaikan task
	err = h.usecase.FinishTask(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	// Beri response sukses
	h.response.JSON(
		w,
		"penyelesaian task sedang di proses",
		nil,
		nil,
	)
}

// GetTaskList menangani request untuk mendapatkan daftar task pengguna
func (h *TaskHandler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	// Ekstrak token dari header Authorization
	tokenString, err := h.extractBearerToken(r)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Verifikasi token JWT
	dataClaims, err := helper.VerifyToken(tokenString)
	if err != nil {
		// Tangani error jika token tidak valid atau kadaluarsa
		if err == jwt.ErrSignatureInvalid || strings.Contains(err.Error(), "token is expired") {
			h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("token expired")))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	// Inisialisasi DTO untuk mendapatkan task
	getDTO := dto.GetTaskReqDTO{
		UserID: dataClaims.UserID, // Ambil UserID dari token
	}

	// Panggil use case untuk mendapatkan daftar task
	resp, err := h.usecase.GetTaskList(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	// Beri response sukses dengan daftar task
	h.response.JSON(
		w,
		"get data task sukses",
		resp,
		nil,
	)
}
