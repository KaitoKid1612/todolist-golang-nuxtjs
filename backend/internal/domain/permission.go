package domain

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"` // Tên quyền
	Description string // Mô tả về quyền
}