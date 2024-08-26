package routes

import (
	"database/sql"
	"net/http"

	"github.com/ghulammuzz/go-restful-template/internal/handler"
	"github.com/ghulammuzz/go-restful-template/internal/repository"
	"github.com/ghulammuzz/go-restful-template/internal/service"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	mux.HandleFunc("/register", userHandler.Register)
}
