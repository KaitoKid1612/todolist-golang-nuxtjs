package models

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"` // Tên dự án
	Description string // Mô tả về dự án

	// Quan hệ với Task (One-to-Many, dự án có thể có nhiều công việc)
	Tasks []Task `gorm:"foreignKey:ProjectID"`
}
