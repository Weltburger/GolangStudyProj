package controllers

import (
	"context"
	"forRoma/pkg/models"
	"forRoma/pkg/statuserror"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type CommentController struct {
	controller *Controller
}

func (commentController *CommentController) CreateComment(c *gin.Context) {
	param := struct {
		Text      string                `json:"text" binding:"required"`
	}{}

	if err := c.Bind(&param); err != nil {
		panic(statuserror.New(404, statuserror.StatusNotFilled, err))
	}

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

	comment, err := commentController.controller.store.CommentRepository().InsertComment(ctx, &models.Comment{
		User:      user,
		Text:      param.Text,
	}, article)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "comment created",
		"value": comment.UUID,
	})
}

func (commentController *CommentController) CommentLikeUnlike(c *gin.Context) {

	value1, ok := c.Get("user")
	if !ok {
		panic(statuserror.NullUser)
	}
	user := value1.(*models.User)

	value2, ok := c.Get("comment")
	if !ok {
		panic(statuserror.NullUser)
	}
	comment := value2.(*models.Comment)

	ctx := context.Background()

	like := &models.LikeComment{
		User: user,
		Comment: comment,
	}
	isExist, err := commentController.controller.store.CommentRepository().LikeExist(ctx, like)
	if err != nil {
		panic(err)
	}
	if isExist {
		if err := commentController.controller.store.CommentRepository().DeleteLike(ctx, like); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "comment has been liked",
		})
	} else {
		if _, err := commentController.controller.store.CommentRepository().InsertLike(ctx, like); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "comment has been unliked",
		})
	}
}

func (commentController *CommentController) ReadComment(c *gin.Context) {
	ctx := context.Background()
	commentUUID, err := uuid.FromString(c.Param("comment"))
	if  err != nil {
		panic(statuserror.New(404, "bad comment uuid", err))
	}

	comment, err := commentController.controller.store.CommentRepository().CommentByUUID(ctx, &models.Comment{
		UUID: commentUUID,
	})

	if err != nil {
		panic(statuserror.New(404, "bad comment uuid", err))
	}

	c.Set("comment", comment)
}
