package author

import (
	"first_project/internal/domain/author/entity"
	"first_project/internal/domain/author/model"
	"github.com/gin-gonic/gin"
)

type Usecase interface {
	CreateAuthor(ctx *gin.Context, authorDTO model.AuthorDTO) error
	GetAuthors(ctx *gin.Context) ([]model.AuthorDTO, error)
}

type Repository interface {
	CreateAuthor(ctx *gin.Context, author entity.Author) error
	GetAuthors(ctx *gin.Context) ([]entity.Author, error)
}
