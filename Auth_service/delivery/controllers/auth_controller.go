package controllers

import (
	"auth-service/domain"

	"github.com/gin-gonic/gin"
)
type UserController struct {
	UserUsecase domain.UserUsecase
}


func NewAuthController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}


func(uc *UserController) Register(c *gin.Context) {
    var user domain.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err := uc.UserUsecase.Register(user)
    if err.Message != "" {
        c.JSON(400, gin.H{
            "status": err.StatusCode,
            "message": err.Message,
        })
        return
    }
    c.JSON(200, gin.H{
        "status": 200,
        "message": "Account created successfully",
    })
}

func(uc *UserController) Login(c *gin.Context) {
    var user domain.User
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result,err := uc.UserUsecase.Login(user.Email, user.Password)
    if err.Message != "" {
        c.JSON(400, gin.H{
            "status": err.StatusCode,
            "message": err.Message,
    })
        return
    }

    c.JSON(200, gin.H{
    "status": 200,        
    "data": result,
})
}