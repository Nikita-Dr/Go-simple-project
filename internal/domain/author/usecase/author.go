package usecase

import (
	"first_project/internal/domain/author"
	"first_project/internal/domain/author/entity"
	"first_project/internal/domain/author/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthorUseCase struct {
	authorRepo author.Repository
}

func NewAuthorUseCase(authorRepo author.Repository) *AuthorUseCase {
	return &AuthorUseCase{
		authorRepo: authorRepo,
	}
}

func (c *AuthorUseCase) CreateAuthor(ctx *gin.Context, authorDTO model.AuthorDTO) error {
	author := entity.NewAuthorUponCreation(authorDTO.Name)
	if err := c.authorRepo.CreateAuthor(ctx, author); err != nil {
		return fmt.Errorf("AuthorUseCase - CreateAuthor - CreateAuthor: %w", err)
	}
	return fmt.Errorf("")
}

func (c *AuthorUseCase) GetAuthors(ctx *gin.Context) ([]model.AuthorDTO, error) {
	authorList, err := c.authorRepo.GetAuthors(ctx)
	if err != nil {
		return []model.AuthorDTO{}, fmt.Errorf("AuthorUseCase - GetAuthors - authorRepo.GetAuthors: %w", err)
	}
	return model.GetAuthorListDTO(authorList), nil
}
