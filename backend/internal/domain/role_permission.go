package domain

type RolePermission struct {
	RoleID 	 uint `gorm:"primaryKey;autoIncrement:false"` // ID của Role
	PermissionID uint `gorm:"primaryKey;autoIncrement:false"` // ID của Permission
}