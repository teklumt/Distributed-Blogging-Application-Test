package usecase

import "blog-service/domain"

type BlogUseCase struct {
	BlogRepository domain.BlogRepository
	
}

func NewBlogUsecase(blogRepository domain.BlogRepository) domain.BlogUsecase {
	return &BlogUseCase{BlogRepository: blogRepository}
}

func (uc *BlogUseCase) CreateBlog(blog domain.Blog) (domain.Blog, domain.ErrorResponse) {
	result, err := uc.BlogRepository.CreateBlog(blog)
	if err != nil {
		return domain.Blog{}, domain.ErrorResponse{Message: "Failed to create blog", StatusCode: 500}
	}

	return result, domain.ErrorResponse{}
}



func (uc *BlogUseCase) GetBlogByID(id uint) (domain.Blog, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetBlogByID(id)
	if err != nil {
		return domain.Blog{}, domain.ErrorResponse{Message: "Blog not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) GetAllBlog() ([]domain.Blog, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetAllBlog()
	if err != nil {
		return []domain.Blog{}, domain.ErrorResponse{Message: "Blog not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) UpdateBlog(blog domain.Blog) (domain.Blog, domain.ErrorResponse) {
	result, err := uc.BlogRepository.UpdateBlog(blog)
	if err != nil {
		return domain.Blog{}, domain.ErrorResponse{Message: "Failed to update blog", StatusCode: 500}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) DeleteBlog(id uint) domain.ErrorResponse {
	err := uc.BlogRepository.DeleteBlog(id)
	if err != nil {
		return domain.ErrorResponse{Message: "Failed to delete blog", StatusCode: 500}
	}

	return domain.ErrorResponse{}
}

func (uc *BlogUseCase) GetBlogByUserID(id uint) ([]domain.Blog, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetBlogByUserID(id)
	if err != nil {
		return []domain.Blog{}, domain.ErrorResponse{Message: "Blog not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}


func (uc *BlogUseCase) CreateComment(comment domain.Comment) (domain.Comment, domain.ErrorResponse) {
	result, err := uc.BlogRepository.CreateComment(comment)
	if err != nil {
		return domain.Comment{}, domain.ErrorResponse{Message: "Failed to create comment", StatusCode: 500}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) GetComentsByPostId(id uint) ([]domain.Comment, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetComentsByPostId(id)
	if err != nil {
		return []domain.Comment{}, domain.ErrorResponse{Message: "Comment not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) GetCommentByID(id uint) (domain.Comment, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetCommentByID(id)
	if err != nil {
		return domain.Comment{}, domain.ErrorResponse{Message: "Comment not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) GetCommentByUserID(id uint) ([]domain.Comment, domain.ErrorResponse) {
	result, err := uc.BlogRepository.GetCommentByUserID(id)
	if err != nil {
		return []domain.Comment{}, domain.ErrorResponse{Message: "Comment not found", StatusCode: 404}
	}

	return result, domain.ErrorResponse{}
}

func (uc *BlogUseCase) DeleteComment(id uint) domain.ErrorResponse {
	err := uc.BlogRepository.DeleteComment(id)
	if err != nil {
		return domain.ErrorResponse{Message: "Failed to delete comment", StatusCode: 500}
	}

	return domain.ErrorResponse{}
}
