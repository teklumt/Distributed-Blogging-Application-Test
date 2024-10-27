package domain

import "time"


type Blog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"user_id" json:"user_id"`
	Title     string    `gorm:"unique;not null" json:"title"`
	Content   string    `gorm:"unique;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt"`

}

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"post_id" json:"post_id"`
	UserID    uint      `gorm:"user_id" json:"user_id"`
	Content   string    `gorm:"unique;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (Blog) TableName() string {
	return "Posts"
}

func (Comment) TableName() string {
	return "Comments"
}




type BlogRepository interface {
	CreateBlog(blog Blog) (Blog, error)
	GetBlogByID(id uint) (Blog, error)
	GetAllBlog() ([]Blog, error)
	UpdateBlog(blog Blog) (Blog, error)
	DeleteBlog(id uint) error

	GetBlogByUserID(id uint) ([]Blog, error)

	CreateComment( comment Comment) (Comment , error)
	GetComentsByPostId(id uint) ([]Comment, error)
	GetCommentByID(id uint) (Comment, error)
	GetCommentByUserID(id uint) ([]Comment, error)
	DeleteComment(id uint) error

}


type BlogUsecase interface {
	CreateBlog(blog Blog) (Blog, ErrorResponse)
	GetBlogByID(id uint) (Blog, ErrorResponse)
	GetAllBlog() ([]Blog, ErrorResponse)
	UpdateBlog(blog Blog) (Blog, ErrorResponse)
	DeleteBlog(id uint) ErrorResponse

	GetBlogByUserID(id uint) ([]Blog, ErrorResponse)

	CreateComment(comment Comment) (Comment, ErrorResponse)
	GetComentsByPostId(id uint) ([]Comment, ErrorResponse)
	GetCommentByID(id uint) (Comment, ErrorResponse)
	GetCommentByUserID(id uint) ([]Comment, ErrorResponse)
	DeleteComment(id uint) ErrorResponse
}