package domain

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"` // Tên vai trò (ví dụ: Admin, User)
	Description string // Mô tả vai trò
}
