package controllers

import (
	"encoding/json"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/domain"
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/messaging"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Publisher   *messaging.Publisher
}

func NewBlogController(blogUsecase domain.BlogUsecase, publisher *messaging.Publisher) *BlogController {
	return &BlogController{
		BlogUsecase: blogUsecase,
		Publisher:   publisher,
	}
}

// @Summary Create Blog
// @Description Create a new blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body domain.Blog true "Blog data"
// @Success 200 {object} domain.Blog
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/create [post]
func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	userIDStr := c.GetString("user_id")
	username := c.GetString("username")

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
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}

	event := messaging.PostCreatedEvent{
		UserID:  uint(userID),
		Message: "New post created by user " + username,
	}
	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal registration event: %v", err)
		return
	}

	err = bc.Publisher.PublishMessage("notification_queue", message)
	if err != nil {
		log.Printf("Failed to publish registration event: %v", err)
	}

	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Get Blog by ID
// @Description Retrieve a blog post by its ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} domain.Blog
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/get/{id} [get]
func (bc *BlogController) GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}

	result, customError := bc.BlogUsecase.GetBlogByID(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Get All Blogs
// @Description Retrieve all blog posts
// @Tags Blog
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Blog
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/get [get]
func (bc *BlogController) GetAllBlog(c *gin.Context) {
	result, customError := bc.BlogUsecase.GetAllBlog()
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Update Blog
// @Description Update an existing blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body domain.Blog true "Updated Blog Data"
// @Success 200 {object} domain.Blog
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/update [put]
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	var blog domain.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, customError := bc.BlogUsecase.UpdateBlog(blog)
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Delete Blog
// @Description Delete a blog post by its ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {string} string "Blog deleted successfully"
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/delete/{id} [delete]
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}
	customError := bc.BlogUsecase.DeleteBlog(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Blog deleted successfully"})
}

// @Summary Get Blog by User ID
// @Description Retrieve all blog posts by a specific user
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} []domain.Blog
// @Failure 400 {object} domain.ErrorResponse
// @Router /blog/get/user/{id} [get]
func (bc *BlogController) GetBlogByUserID(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetBlogByUserID(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Create Comment
// @Description Create a comment for a blog post
// @Tags Comment
// @Accept json
// @Produce json
// @Param comment body domain.Comment true "Comment data"
// @Success 200 {object} domain.Comment
// @Failure 400 {object} domain.ErrorResponse
// @Router /comment/create [post]
func (bc *BlogController) CreateComment(c *gin.Context) {
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
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}

	event := messaging.CommentPostedEvent{
		UserID:  uint(userID),
		Message: "New comment posted by user " + userIDStr,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal registration event: %v", err)
		return
	}

	err = bc.Publisher.PublishMessage("notification_queue", message)
	if err != nil {
		log.Printf("Failed to publish registration event: %v", err)
	}

	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Get Comments by Post ID
// @Description Retrieve all comments for a specific blog post
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} []domain.Comment
// @Failure 400 {object} domain.ErrorResponse
// @Router /comment/get/{id} [get]
func (bc *BlogController) GetComentsByPostId(c *gin.Context) {
	id := c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}
	result, customError := bc.BlogUsecase.GetComentsByPostId(uint(blogID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Get Comment by ID
// @Description Retrieve a single comment by its ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} domain.Comment
// @Failure 400 {object} domain.ErrorResponse
// @Router /comment/get/comment/{id} [get]
func (bc *BlogController) GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}

	result, customError := bc.BlogUsecase.GetCommentByID(uint(commentID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": result})
}

// @Summary Delete Comment
// @Description Delete a comment by its ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {string} string "Comment deleted successfully"
// @Failure 400 {object} domain.ErrorResponse
// @Router /comment/delete/{id} [delete]
func (bc *BlogController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment ID"})
		return
	}

	customError := bc.BlogUsecase.DeleteComment(uint(commentID))
	if customError.Message != "" {
		c.JSON(400, gin.H{"status": customError.StatusCode, "message": customError.Message})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Comment deleted successfully"})
}


// @Summary Get Comments by User ID
// @Description Retrieve all comments made by a specific user
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} []domain.Comment
// @Failure 400 {object} domain.ErrorResponse
// @Router /comment/get/user/{id} [get]
func (bc *BlogController) GetCommentByUserID(c *gin.Context) {
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
	  "data":   result,
	})
  }
  