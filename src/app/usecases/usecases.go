package usecases

import (
	taskUC "todo_list/src/app/usecases/task"
	userUC "todo_list/src/app/usecases/user"
)

type AllUseCases struct {
	UserUC userUC.UserUCInterface
	TaskUC taskUC.TaskUCInterface
}
