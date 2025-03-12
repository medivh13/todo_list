package route

import (
	"net/http"

	handlers "todo_list/src/interface/rest/handler/user"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func UserRouter(h handlers.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.RegisterUser)
	r.Post("/sign-in", h.SignIn)
	r.Post("/refresh-token", h.RefreshToken)

	return r
}