package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kakuzops/api-code-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.FingById)
	r.GET("/user", controller.GetAllUser)
	r.POST("/user/create", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/delete/:userId", controller.DeleteUser)
}
