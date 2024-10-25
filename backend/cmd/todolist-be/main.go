package main

import (
	"backend/api/handlers"
	"backend/api/routes"
	"backend/api/services"
	"backend/config"
	"backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo Gin router
	r := gin.Default()

	// Kết nối tới database
	db := config.SetupDB()

	// Khởi tạo repository
	roleRepo := repository.NewRoleRepository(db)

	// Khởi tạo service
	roleService := services.NewRoleService(roleRepo)

	// Khởi tạo handler
	roleHandler := handlers.NewRoleHandler(roleService)

	// Đăng ký các route
	routes.RegisterRoutes(r, roleHandler)

	// Chạy server trên cổng 8080
	r.Run(":8080")
}