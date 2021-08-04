package controllers

import (
	"context"
	"forRoma/pkg/models"
	"forRoma/pkg/statuserror"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type ArticleController struct {
	controller *Controller
}

func (articleController *ArticleController) CreateArticle(c *gin.Context) {
	param := struct {
		Title     string                `json:"title" binding:"required"`
		Text      string                `json:"text" binding:"required"`
	}{}

	if err := c.Bind(&param); err != nil {
		panic(statuserror.New(404, statuserror.StatusNotFilled, err))
	}

	value, ok := c.Get("user")
	if !ok {
		panic(statuserror.NullUser)
	}

	user := value.(*models.User)

	ctx := context.Background()

	article, err := articleController.controller.store.ArticleRepository().InsertArticle(ctx, &models.Article{
		User:      user,
		Title:     param.Title,
		Text:      param.Text,
	})

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "article created",
		"value": article.UUID,
	})
}

func (articleController *ArticleController) ReadArticle(c *gin.Context) {
	ctx := context.Background()
	articleUUID, err := uuid.FromString(c.Param("article"))
	if  err != nil {
		panic(statuserror.New(404, "bad article uuid", err))
	}

	article, err := articleController.controller.store.ArticleRepository().ArticleByUUID(ctx, &models.Article{
		UUID: articleUUID,
	})

	if err != nil {
		panic(statuserror.New(404, "bad article uuid", err))
	}

	c.Set("article", article)
}

func (articleController *ArticleController) ArticleLikeUnlike(c *gin.Context) {

	value1, ok := c.Get("user")
	if !ok {
		panic(statuserror.NullUser)
	}
	user := value1.(*models.User)

	value2, ok := c.Get("article")
	if !ok {
		panic(statuserror.NullUser)
	}
	article := value2.(*models.Article)

	ctx := context.Background()

	like := &models.LikeArticle{
		User: user,
		Article: article,
	}
	isExist, err := articleController.controller.store.ArticleRepository().LikeExist(ctx, like)
	if err != nil {
		panic(err)
	}
	if isExist {
		if err := articleController.controller.store.ArticleRepository().DeleteLike(ctx, like); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "article has been liked",
		})
	} else {
		if _, err := articleController.controller.store.ArticleRepository().InsertLike(ctx, like); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "article has been unliked",
		})
	}
}

func (articleController *ArticleController) GetAll(c *gin.Context) {

	param := struct {
		Take int64 `form:"take,default=10"`
		Skip int64 `form:"skip,default=0"`
	}{}

	if err := c.Bind(&param); err != nil {
		panic(err)
	}

	ctx := context.Background()

	articles, total, err := articleController.controller.store.ArticleRepository().AllArticles(ctx, param.Take, param.Skip)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": articles,
		"metadata": gin.H{
			"take": param.Take,
			"skip": param.Skip,
			"total": total,
		},
	})
}
