package routers

import (
	"blog-service/delivery/controllers"
	"blog-service/infrastructure"
	"blog-service/messaging"
	"blog-service/repository"
	"blog-service/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, publisher *messaging.Publisher) {

    BlogRepo := repository.NewBlogRepository()
	UserUsecase := usecase.NewBlogUsecase(BlogRepo)
	
    BlogController := controllers.NewBlogController(UserUsecase, publisher)

    BlogGroup := r.Group("/blog")
	BlogGroup.Use(infrastructure.AuthMiddleware())

    {
		BlogGroup.POST("/create", BlogController.CreateBlog)
		BlogGroup.GET("/get", BlogController.GetAllBlog)
		BlogGroup.GET("/get/:id", BlogController.GetBlogByID)
		BlogGroup.PUT("/update", BlogController.UpdateBlog)
		BlogGroup.DELETE("/delete/:id", BlogController.DeleteBlog)
		BlogGroup.GET("/get/user/:id", BlogController.GetBlogByUserID)
		
        
    }

	CommentGroup := r.Group("/comment")
	CommentGroup.Use(infrastructure.AuthMiddleware())
	{
		CommentGroup.POST("/create", BlogController.CreateComment)
		CommentGroup.GET("/get/:id", BlogController.GetComentsByPostId)
		CommentGroup.GET("/get/user/:id", BlogController.GetCommentByUserID)
		CommentGroup.GET("/get/comment/:id", BlogController.GetCommentByID)	
		CommentGroup.DELETE("/delete/:id", BlogController.DeleteComment)


	}
}