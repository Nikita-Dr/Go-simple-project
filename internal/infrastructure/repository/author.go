package repository

import (
	"first_project/internal/domain/author/entity"
	"first_project/pkg/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthorRepository struct {
	db *postgres.Postrgress
}

func NewAuthorRepository(db *postgres.Postrgress) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) CreateAuthor(ctx *gin.Context, author entity.Author) error {
	if err := r.db.NewInsert().Model(&author).Scan(ctx); err != nil {
		return fmt.Errorf("AuthorRepository - CreateAuthor - NewInsert: %w", err)
	}
	return nil
}

func (r *AuthorRepository) GetAuthors(ctx *gin.Context) ([]entity.Author, error) {
	authorList := []entity.Author{}
	if err := r.db.NewSelect().Model(&authorList).Scan(ctx); err != nil {
		return []entity.Author{}, fmt.Errorf("AuthorRepository - GetAuthors - NewSelect: %w", err)
	}
	return authorList, nil
}
