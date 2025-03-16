package task

import (
	"log"
	dto "todo_list/src/app/dto/task"

	"github.com/jmoiron/sqlx"
)

// TaskRepository mendefinisikan metode yang harus diimplementasikan

type TaskRepository interface {
	GetTaskList(req *dto.GetTaskReqDTO) ([]*dto.GetTaskRespDTO, error)
}

// Query SQL untuk berbagai operasi database
const (
	GetTaskList = `SELECT id, title, status, expires_at from public.tasks where user_id = $1`
)

// Struct untuk menyimpan statement yang telah diprepare
var statement PreparedStatement

type PreparedStatement struct {
	getTaskList *sqlx.Stmt
}

type taskRepo struct {
	Connection *sqlx.DB
}

// NewUserRepository menginisialisasi UserRepo dan menyiapkan prepared statement
func NewTaskRepository(db *sqlx.DB) TaskRepository {
	repo := &taskRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

// Preparex menyiapkan statement SQL yang telah diprepare
func (p *taskRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

// InitPreparedStatement menginisialisasi prepared statement untuk query tertentu
func InitPreparedStatement(m *taskRepo) {
	statement = PreparedStatement{
		getTaskList: m.Preparex(GetTaskList),
	}
}

func (repo *taskRepo) GetTaskList(req *dto.GetTaskReqDTO) ([]*dto.GetTaskRespDTO, error) {
	var resp []*dto.GetTaskRespDTO
	err := statement.getTaskList.Select(&resp, req.UserID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
