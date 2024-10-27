package controllers

import (
	"auth-service/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)
type UserController struct {
	UserUsecase domain.UserUsecase
}


func NewAuthController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}


func(uc *UserController) GerUserByID(c *gin.Context) {

    
    id := c.Param("id")

    userID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid user ID"})
        return
    }

    result, customError := uc.UserUsecase.GerUserByID(uint(userID))
    if customError.Message != "" {
        c.JSON(400, gin.H{
            "status": customError.StatusCode,
            "message": customError.Message,
        })
        return
    }
    c.JSON(200, gin.H{
        "status": 200,
        "data": result,
    })
}

func(uc *UserController) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user domain.User


    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Convert id from string to uint
    userID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid user ID"})
        return
    }

    result, customError := uc.UserUsecase.UpdateUser(uint(userID), user)
    if customError.Message != "" {
        c.JSON(400, gin.H{
            "status": customError.StatusCode,
            "message": customError.Message,
        })
        return
    }

    c.JSON(200, gin.H{
        "status": 200,
        "data": result,
    })
}