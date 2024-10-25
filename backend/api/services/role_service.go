package services

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
)

type RoleService struct {
	roleRepo *repository.RoleRepository
}

// Khởi tạo RoleService mới
func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

// Lấy tất cả các Role
func (s *RoleService) GetAllRoles() ([]domain.Role, error) {
	return s.roleRepo.GetAllRoles()
}

// Tạo một Role mới
func (s *RoleService) CreateRole(role *domain.Role) error {
	// Business logic: Ví dụ, kiểm tra điều kiện nếu cần
	if role.Name == "" {
		return errors.New("Role name cannot be empty")
	}

	// Thực hiện lưu role vào cơ sở dữ liệu
	return s.roleRepo.CreateRole(role)
}

// Xóa Role theo ID
func (s *RoleService) DeleteRole(id uint) error {
	return s.roleRepo.DeleteRole(id)
}
