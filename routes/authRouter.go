package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/willyblord/goAuth/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("signup", controllers.Signup())
	incomingRoutes.POST("signin", controllers.Signin())
}
