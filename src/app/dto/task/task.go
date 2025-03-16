package task

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateTaskReqDTO digunakan untuk membuat task baru
type CreateTaskReqDTO struct {
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (dto *CreateTaskReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Title, validation.Required),
		validation.Field(&dto.ExpiresAt, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type FinishtTaskReqDTO struct {
	ID int64 `json:"id"`
}

// UpdateTaskReqDTO digunakan untuk memperbarui task yang sudah ada
type GetTaskReqDTO struct {
	UserID int64 `json:"id"`
}

type ExpireTaskReqDTO struct {
	ID int64 `json:"id"`
}

type GetTaskRespDTO struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Status    string    `json:"status" db:"status"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}
