package routes

import (
	"backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(router *gin.Engine, roleHandler *handlers.RoleHandler) {
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.GET("/", roleHandler.GetRoles)
		roleRoutes.POST("/", roleHandler.CreateRole)
		roleRoutes.DELETE("/:id", roleHandler.DeleteRole)
	}
}

func RegisterRoutes(router *gin.Engine, roleHandler *handlers.RoleHandler) {
	RegisterRoleRoutes(router, roleHandler)
}
