package server

import (
	"context"
	"forRoma/pkg/config"
	"forRoma/pkg/controllers"
	"forRoma/pkg/statuserror"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	sync.Once
}

func routing(conf *config.Config) (*gin.Engine, error) {
	controller, err := controllers.New(conf)
	if err != nil {
		return nil, err
	}

	router := gin.New()

	recovery := router.Group("/", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case statuserror.IStatusError:
					c.JSON(e.HttpCode(), gin.H{
						"code":    e.StatusCode(),
						"message": e.Error(),
					})
				case error:
					c.JSON(500, gin.H{
						"error": e.Error(),
					})
				case string:
					c.JSON(500, gin.H{
						"error": e,
					})
				default:
					c.JSON(500, gin.H{
						"error": "undefined error",
					})
				}
				c.Abort()
			}
		}()
		c.Next()
	})

	recovery.POST("/login", controller.UserController().Login)
	recovery.POST("/registration", controller.UserController().Registration)

	sessionGroup := recovery.Group("/", controller.SessionController().CheckAuthorisation)

	sessionGroup.GET("/ok", controller.UserController().OK)

	sessionGroup.GET("/articles", controller.ArticleController().GetAll)

	sessionGroup.POST("/create/article", controller.ArticleController().CreateArticle)
	sessionGroup.POST("/article/:article/create/comment",
		controller.ArticleController().ReadArticle,
		controller.CommentController().CreateComment)

	sessionGroup.POST("/article/:article/like_unlike",
		controller.ArticleController().ReadArticle,
		controller.ArticleController().ArticleLikeUnlike)

	sessionGroup.POST("/article/:article/:comment/like_unlike",
		controller.ArticleController().ReadArticle,
		controller.CommentController().ReadComment,
		controller.CommentController().CommentLikeUnlike)

	return router, nil
}

func New(path string) (*Server, error) {
	conf, err := config.New(path)
	if err != nil {
		return nil, err
	}

	router, err := routing(conf)
	if err != nil {
		return nil, err
	}

	return &Server{config: conf, router: router}, nil
}

func NewTest(path string) (*gin.Engine, error) {
	conf, err := config.New(path)
	if err != nil {
		return nil, err
	}

	router, err := routing(conf)
	if err != nil {
		return nil, err
	}

	return router, nil
}

func (server *Server) Start() error {
	serv := &http.Server{
		Addr:    server.config.Addr,
		Handler: server.router,
	}
	serv.SetKeepAlivesEnabled(true)
	ch := make(chan error, 1)
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			ch <- err
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
		case err := <-ch:
			return err
		case <-interrupt:
	}

	timeout, CancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer CancelFunc()

	if err := serv.Shutdown(timeout); err != nil {
		return err
	}

	return nil
}
