package model

import "first_project/internal/domain/author/entity"

type AuthorDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (AuthorDTO) TableName() string {
	return "authors"
}

func GetAuthorDTO(author entity.Author) AuthorDTO {
	return AuthorDTO{
		ID:   author.ID,
		Name: author.Name,
	}
}

func GetAuthorListDTO(authorList []entity.Author) []AuthorDTO {
	authorListDTO := []AuthorDTO{}
	for _, author := range authorList {
		authorListDTO = append(authorListDTO, GetAuthorDTO(author))
	}
	return authorListDTO
}
