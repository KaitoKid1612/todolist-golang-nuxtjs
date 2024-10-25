package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"todolist-backend/database"
)

func main() {
	r := gin.Default()

	// Kết nối cơ sở dữ liệu
	database.SetupDB() // Gọi hàm SetupDB từ package database

	// Lấy đối tượng *sql.DB để kiểm tra kết nối
	sqlDB, err := database.DB.DB() // Sử dụng biến DB toàn cục từ package database
	if err != nil {
		log.Fatalf("Failed to get db connection: %v", err)
	}
	defer sqlDB.Close() // Đảm bảo đóng kết nối khi không sử dụng

	// Thiết lập các route (sẽ thêm sau)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the To-Do List API"})
	})

	// Khởi động server
	r.Run(":8080") // Chạy trên cổng 8080
}
