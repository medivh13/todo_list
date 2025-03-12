package usecases

import (
	userUC "todo_list/src/app/usecases/user"
)

type AllUseCases struct {
	UserUC userUC.UserUCInterface
}
