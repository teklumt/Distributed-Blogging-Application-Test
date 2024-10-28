package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/domain"
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/messaging"

	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Publisher *messaging.Publisher
}

func NewBlogController(blogUsecase domain.BlogUsecase, publisher *messaging.Publisher ) *BlogController {
	return &BlogController{
		BlogUsecase: blogUsecase,
		Publisher: publisher,
	}
}



func(bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	userIDStr := c.GetString("user_id")
	username := c.GetString("username")

	fmt.Println("User ID: 9696969696", userIDStr)
	
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	blog.UserID = uint(userID)

	result, customError := bc.BlogUsecase.CreateBlog(blog)
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}

	// Create a structured event

	event := messaging.PostCreatedEvent{
		UserID:   uint(userID),
		Message: "New post created by user " + username,
	}

	message, errr := json.Marshal(event)
    if errr != nil {
        log.Printf("Failed to marshal registration event: %v", err)
        return
    }


	errr = bc.Publisher.PublishMessage("notification_queue", message)

	if errr != nil {
		log.Printf("Failed to publish registration event: %v", err)

	}
		
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}


func(bc *BlogController)GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	result, customError := bc.BlogUsecase.GetBlogByID(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) GetAllBlog(c *gin.Context) {
	result, customError := bc.BlogUsecase.GetAllBlog()
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) UpdateBlog(c *gin.Context) {
	var blog domain.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, customError := bc.BlogUsecase.UpdateBlog(blog)
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}
	customError := bc.BlogUsecase.DeleteBlog(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "Blog deleted successfully",
	})
}


func(bc *BlogController) GetBlogByUserID(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetBlogByUserID(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}



func(bc *BlogController) CreateComment(c *gin.Context) {
	var comment domain.Comment
	userIDStr := c.GetString("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.UserID = uint(userID)

	result, customError := bc.BlogUsecase.CreateComment(comment)
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}


	// Create a structured event

	event := messaging.CommentPostedEvent{
		UserID: uint(userID),
		Message: "New comment posted by user " + userIDStr,
	}

	message, errr := json.Marshal(event)
	if errr != nil {
		log.Printf("Failed to marshal registration event: %v", err)
		return
	}

	errr = bc.Publisher.PublishMessage("notification_queue", message)
	
	if errr != nil {
		log.Printf("Failed to publish registration event: %v", err)
	}




	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) GetComentsByPostId(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetComentsByPostId(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) GetCommentByID(c *gin.Context) {
	id := c.Param("id")

	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetCommentByID(uint(commentID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) GetCommentByUserID(c *gin.Context) {
	id := c.Param("id")

	UserID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetCommentByUserID(uint(UserID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data": result,
	})
}

func(bc *BlogController) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}
	customError := bc.BlogUsecase.DeleteComment(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{
			"status":  customError.StatusCode,
			"message": customError.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "Comment deleted successfully",
	})
}

