package http

import (
	"first_project/internal/domain/author"
	"first_project/internal/domain/author/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authorController struct {
	authorUC author.Usecase
}

func NewAutorConthroller(handler *gin.Engine, authorUC author.Usecase) {
	c := authorController{
		authorUC: authorUC,
	}

	handler.POST("/author", c.createAuthor)
	handler.GET("/author", c.getAuthors)
}

func (c *authorController) createAuthor(ctx *gin.Context) {
	author := model.AuthorDTO{}
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.authorUC.CreateAuthor(ctx, author); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (c *authorController) getAuthors(ctx *gin.Context) {
	authorList, err := c.authorUC.GetAuthors(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, authorList)
}
