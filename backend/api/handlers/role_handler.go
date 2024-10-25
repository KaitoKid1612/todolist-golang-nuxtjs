package handlers

import (
	"backend/internal/domain"
	"backend/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// Lấy danh sách tất cả các role
func (h *RoleHandler) GetRoles(c *gin.Context) {
	println("GetRoles")
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// Tạo role mới
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.roleService.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// Xóa role
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.roleService.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
