package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func setupDB() *gorm.DB {
	dsn := "user=kaito password=123456 dbname=todolist port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func main() {
	r := gin.Default()

	// Kết nối cơ sở dữ liệu
	db := setupDB()

	// Sử dụng db để kiểm tra kết nối (chẳng hạn)
	sqlDB, err := db.DB() // Lấy đối tượng *sql.DB từ GORM
	if err != nil {
		log.Fatalf("Failed to get db connection: %v", err)
	}
	defer sqlDB.Close() // Đảm bảo đóng kết nối khi không còn sử dụng

	// Thiết lập các route (sẽ thêm sau)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the To-Do List API"})
	})

	// Khởi động server
	r.Run(":8080") // Chạy trên cổng 8080
}
