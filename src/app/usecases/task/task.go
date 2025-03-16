package task

import (
	"encoding/json"
	"log"
	dto "todo_list/src/app/dto/task"                   // Import DTO untuk Task
	repo "todo_list/src/app/repositories/task"        // Import repository Task
	natsPublisher "todo_list/src/infra/broker/nats/publisher" // Import publisher NATS
	Const "todo_list/src/infra/constants"             // Import constants
)

// TaskUCInterface mendefinisikan contract untuk Task Use Case
type TaskUCInterface interface {
	AddTask(req *dto.CreateTaskReqDTO) error
	FinishTask(req *dto.FinishtTaskReqDTO) error
	GetTaskList(req *dto.GetTaskReqDTO) ([]*dto.GetTaskRespDTO, error)
}

// taskUseCase adalah implementasi dari TaskUCInterface
type taskUseCase struct {
	Publisher natsPublisher.PublisherInterface // Publisher untuk event NATS
	Repo      repo.TaskRepository              // Repository untuk mengakses database
}

// NewTaskUseCase membuat instance taskUseCase
func NewTaskUseCase(p natsPublisher.PublisherInterface, r repo.TaskRepository) TaskUCInterface {
	return &taskUseCase{
		Publisher: p,
		Repo:      r,
	}
}

// AddTask mengirimkan task baru ke NATS
func (uc *taskUseCase) AddTask(req *dto.CreateTaskReqDTO) error {
	newData, _ := json.Marshal(req)               // Serialize request ke JSON
	err := uc.Publisher.Nats(newData, Const.ADD_TASK) // Kirim ke NATS
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// FinishTask mengirimkan event selesai task ke NATS
func (uc *taskUseCase) FinishTask(req *dto.FinishtTaskReqDTO) error {
	newData, _ := json.Marshal(req)                  // Serialize request ke JSON
	err := uc.Publisher.Nats(newData, Const.FINISH_TASK) // Kirim ke NATS
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetTaskList mengambil daftar task dari repository
func (uc *taskUseCase) GetTaskList(req *dto.GetTaskReqDTO) ([]*dto.GetTaskRespDTO, error) {
	resp, err := uc.Repo.GetTaskList(req) // Ambil data task dari repository
	if err != nil {
		return nil, err
	}
	return resp, nil
}
