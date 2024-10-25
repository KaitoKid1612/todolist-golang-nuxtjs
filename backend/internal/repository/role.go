package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

// RoleRepository là nơi giao tiếp với database
type RoleRepository struct {
	DB *gorm.DB
}

// Khởi tạo RoleRepository mới
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

// Lấy tất cả roles
func (r *RoleRepository) GetAllRoles() ([]domain.Role, error) {
	var roles []domain.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Tạo role mới
func (r *RoleRepository) CreateRole(role *domain.Role) error {
	return r.DB.Create(role).Error
}

// Xóa role theo ID
func (r *RoleRepository) DeleteRole(id uint) error {
	return r.DB.Delete(&domain.Role{}, id).Error
}
