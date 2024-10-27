package repository

import (
	"auth-service/config"
	"auth-service/domain"
	"auth-service/messaging"
)

type BlogRepository struct{}

// Initialize a new user repository
func NewBlogRepository() domain.BlogRepository {
	return &BlogRepository{}
}


func (br *BlogRepository) CreateBlog(blog domain.Blog) (domain.Blog, error) {
	err := config.DB.Create(&blog).Error


	// Publish new blog posted event
	event := messaging.NewBlogPostedEvent{
		UserID: "all",
		Message: "New blog posted",
	}
	
	err = messaging.PublishNewBlogPostedEvent(event)
	if err != nil {
		return domain.Blog{}, err
	}


	return blog, err
}

func (br *BlogRepository) GetBlogByID(id uint) (domain.Blog, error) {
	var blog domain.Blog
	err := config.DB.Where("id = ?", id).First(&blog).Error
	return blog, err
}

func (br *BlogRepository) GetAllBlog() ([]domain.Blog, error) {
	var blogs []domain.Blog
	err := config.DB.Find(&blogs).Error
	return blogs, err
}

func (br *BlogRepository) UpdateBlog(blog domain.Blog) (domain.Blog, error) {
	err := config.DB.Model(&domain.Blog{}).Where("id = ?", blog.ID).Updates(blog).Error
	return blog, err
}

func (br *BlogRepository) DeleteBlog(id uint) error {
	err := config.DB.Where("id = ?", id).Delete(&domain.Blog{}).Error
	return err
}

func (br *BlogRepository) GetBlogByUserID(id uint) ([]domain.Blog, error) {
	var blogs []domain.Blog
	err := config.DB.Where("user_id = ?", id).Find(&blogs).Error
	return blogs, err
}


func (br *BlogRepository) CreateComment(comment domain.Comment) (domain.Comment, error) {
	err := config.DB.Create(&comment).Error

	// Publish new blog posted event

	event := messaging.CommentPostedEvent{
		PostID: comment.PostID,
		Message: "New comment posted",
	}

	err = messaging.PublishCommentPostedEvent(event)
	if err != nil {
		return domain.Comment{}, err
	}

	return comment, err
}

func (br *BlogRepository) GetComentsByPostId(id uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := config.DB.Where("post_id = ?", id).Find(&comments).Error
	return comments, err
}

func (br *BlogRepository) GetCommentByID(id uint) (domain.Comment, error) {
	var comment domain.Comment
	err := config.DB.Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (br *BlogRepository) GetCommentByUserID(id uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := config.DB.Where("user_id = ?", id).Find(&comments).Error
	return comments, err
}

func (br *BlogRepository) DeleteComment(id uint) error {
	err := config.DB.Where("id = ?", id).Delete(&domain.Comment{}).Error
	return err
}
