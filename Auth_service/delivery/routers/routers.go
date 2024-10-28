package routers

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/delivery/controllers"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/infrastructure"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/repository"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/usecase"

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
    }
}