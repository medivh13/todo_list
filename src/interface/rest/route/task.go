package route

import (
	"net/http"

	handlers "todo_list/src/interface/rest/handler/task"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func TaskRouter(h handlers.TaskHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.AddTask)
	r.Patch("/", h.FinishTask)
	r.Get("/", h.GetTaskList)

	return r
}
