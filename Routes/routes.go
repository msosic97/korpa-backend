// sve rute , inicijalizacija servera,
package routes

import (
	"net/http"

	routes "github.com/msosic97/korpa-backend/routes/user"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/user/register", routes.RegisterUserHandler)

	return router
}
