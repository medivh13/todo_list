package task

import (
	dto "todo_list/src/app/dto/task"
	repo "todo_list/src/app/repositories/task"

	"github.com/stretchr/testify/mock"
)

type MockTask struct {
	mock.Mock
}

func NewMockTask() *MockTask {
	return &MockTask{}
}

var _ repo.TaskRepository = &MockTask{}

func (o *MockTask) GetTaskList(req *dto.GetTaskReqDTO) ([]*dto.GetTaskRespDTO, error) {
	args := o.Called(req)

	var (
		resp []*dto.GetTaskRespDTO
		err  error
	)

	if n, ok := args.Get(0).([]*dto.GetTaskRespDTO); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}
