package task

import (
	"encoding/json"
	"errors"
	"time"
	mockPubliser "todo_list/mock/infra/broker/nats/publisher"
	mockRepo "todo_list/mock/repositories/task"

	"testing"
	dto "todo_list/src/app/dto/task"

	Const "todo_list/src/infra/constants"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserUseCase struct {
	mock.Mock
}

type UserUseCaseList struct {
	suite.Suite

	useCase        TaskUCInterface
	mockRepo       *mockRepo.MockTask
	mockPubliser   *mockPubliser.MockPublisher
	dtoAddTask     *dto.CreateTaskReqDTO
	dtoFinishTask  *dto.FinishtTaskReqDTO
	dtoGetTaskList *dto.GetTaskReqDTO
}

func (suite *UserUseCaseList) SetupTest() {

	suite.mockRepo = new(mockRepo.MockTask)
	suite.mockPubliser = new(mockPubliser.MockPublisher)
	suite.useCase = NewTaskUseCase(suite.mockPubliser, suite.mockRepo)

	expiresAt, _ := time.Parse(time.RFC3339, "2025-03-16T12:41:00Z")

	suite.dtoAddTask = &dto.CreateTaskReqDTO{
		UserID:    1,
		Title:     "test",
		ExpiresAt: expiresAt,
	}

	suite.dtoFinishTask = &dto.FinishtTaskReqDTO{
		ID: 1,
	}

	suite.dtoGetTaskList = &dto.GetTaskReqDTO{
		UserID: 1,
	}

}

func (u *UserUseCaseList) TestAddTaskSuccess() {
	newData, _ := json.Marshal(u.dtoAddTask)
	u.mockPubliser.Mock.On("Nats", newData, Const.ADD_TASK).Return(nil)
	err := u.useCase.AddTask(u.dtoAddTask)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestAddTaskFail() {
	newData, _ := json.Marshal(u.dtoAddTask)
	u.mockPubliser.Mock.On("Nats", newData, Const.ADD_TASK).Return(errors.New(mock.Anything))
	err := u.useCase.AddTask(u.dtoAddTask)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestFinistTaskSuccess() {
	newData, _ := json.Marshal(u.dtoFinishTask)
	u.mockPubliser.Mock.On("Nats", newData, Const.FINISH_TASK).Return(nil)
	err := u.useCase.FinishTask(u.dtoFinishTask)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestFinistTaskFail() {
	newData, _ := json.Marshal(u.dtoFinishTask)
	u.mockPubliser.Mock.On("Nats", newData, Const.FINISH_TASK).Return(errors.New(mock.Anything))
	err := u.useCase.FinishTask(u.dtoFinishTask)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestGetTaskListSuccess() {
	u.mockRepo.Mock.On("GetTaskList", u.dtoGetTaskList).Return(mock.Anything, nil)
	_, err := u.useCase.GetTaskList(u.dtoGetTaskList)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestGetTaskListFail() {
	u.mockRepo.Mock.On("GetTaskList", u.dtoGetTaskList).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetTaskList(u.dtoGetTaskList)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UserUseCaseList))
}
