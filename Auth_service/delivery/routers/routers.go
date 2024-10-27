package routers

import (
	"auth-service/delivery/controllers"
	"auth-service/infrastructure"
	"auth-service/repository"
	"auth-service/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

    AuthRepo := repository.NewAuthRepository()
    
	tokenGen := infrastructure.NewTokenGenerator()
	passwordSvc := infrastructure.NewPasswordService()


	AuthUsecase := usecase.NewAuthUsecase(AuthRepo , tokenGen, passwordSvc)
	
    AuthController := controllers.NewAuthController(AuthUsecase)

    authGroup := r.Group("/auth")
    {
        authGroup.POST("/login", AuthController.Login)
        authGroup.POST("/register", AuthController.Register)
        // authGroup.POST("/refresh-token", controllers.RefreshToken)
    }
}