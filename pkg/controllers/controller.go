package controllers

import (
	"forRoma/pkg/config"
	"forRoma/pkg/store"
)

type Controller struct {
	store                 *store.Store
	JWTKey                []byte
	userController        *UserController
	sessionController     *SessionController
	articleController     *ArticleController
	commentController     *CommentController
}

func (controller *Controller) UserController() *UserController {
	if controller.userController != nil {
		return controller.userController
	}

	controller.userController = &UserController{controller: controller}

	return controller.userController
}

func (controller *Controller) ArticleController() *ArticleController {
	if controller.articleController != nil {
		return controller.articleController
	}

	controller.articleController = &ArticleController{controller: controller}

	return controller.articleController
}

func (controller *Controller) CommentController() *CommentController {
	if controller.commentController != nil {
		return controller.commentController
	}

	controller.commentController = &CommentController{controller: controller}

	return controller.commentController
}

func (controller *Controller) SessionController() *SessionController {
	if controller.sessionController != nil {
		return controller.sessionController
	}

	controller.sessionController = &SessionController{controller: controller}

	return controller.sessionController
}


func New(cg *config.Config) (*Controller, error) {

	s, err := store.Open(cg)
	if err != nil {
		return nil, err
	}

	return &Controller{store: s, JWTKey: cg.JWTKey}, nil
}
