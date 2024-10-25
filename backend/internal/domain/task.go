package domain

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"` // Tiêu đề của công việc
	Description string    // Mô tả công việc
	Status      string    `gorm:"default:'pending'"` // Trạng thái công việc (pending, completed)
	CreatedBy   uint      // Người tạo công việc (User ID)
	AssignedTo  uint      // Người được giao công việc
	StartDate   time.Time // Ngày bắt đầu
	DueDate     time.Time // Ngày hết hạn

	// Quan hệ với Project (Many-to-One, công việc thuộc về một dự án)
	ProjectID uint
	Project   Project

	// Quan hệ với User (Many-to-Many, nhiều người dùng có thể nhận task)
	Users []User `gorm:"many2many:task_assignments;"`
}
