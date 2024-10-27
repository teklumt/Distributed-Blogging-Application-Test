package routers

import (
	"auth-service/delivery/controllers"
	"auth-service/infrastructure"
	"auth-service/repository"
	"auth-service/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

    UserRepo := repository.NewAuthRepository()
    
	tokenGen := infrastructure.NewTokenGenerator()
	passwordSvc := infrastructure.NewPasswordService()


	UserUsecase := usecase.NewAuthUsecase(UserRepo , tokenGen, passwordSvc)
	
    UserController := controllers.NewAuthController(UserUsecase)

    UserGroup := r.Group("/user")
    {
		UserGroup.GET("/:id", UserController.GerUserByID)
		UserGroup.PATCH("/:id", UserController.UpdateUser)
        
    }
}