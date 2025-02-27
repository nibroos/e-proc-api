package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/e-proc-api/service/internal/controller/rest"
	"github.com/nibroos/e-proc-api/service/internal/repository"
	"github.com/nibroos/e-proc-api/service/internal/service"
	"gorm.io/gorm"
)

func SetupUserRoutes(users fiber.Router, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	userRepo := repository.NewUserRepository(gormDB, sqlDB)
	userService := service.NewUserService(userRepo)
	customerService := service.NewCustomerService(repository.NewCustomerRepository(gormDB, sqlDB))
	userController := rest.NewUserController(userService, customerService)

	// prefix /users

	users.Post("/index-user", userController.GetUsers)
	users.Post("/show-user", userController.GetUserByID)
	users.Post("/create-user", userController.CreateUser)
	users.Post("/update-user", userController.UpdateUser)
	users.Post("/delete-user", userController.DeleteUser)
	users.Post("/restore-user", userController.RestoreUser)
}
