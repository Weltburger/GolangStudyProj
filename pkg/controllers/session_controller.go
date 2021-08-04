package controllers

import (
	"context"
	"fmt"
	"forRoma/pkg/models"
	"forRoma/pkg/statuserror"
	"github.com/gin-gonic/gin"
	"strings"
)

type SessionController struct {
	controller *Controller
}

func (sessionController *SessionController) CheckAuthorisation(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	key, err := models.Auth(sessionController.controller.JWTKey, reqToken)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	user := &models.User{}
	if err := sessionController.controller.store.Redis.Get(ctx, fmt.Sprintf("access_token:%s", key)).
		Scan(&user.UUID); err != nil {
		panic(statuserror.NotAuthorized)
	}

	user, err = sessionController.controller.store.UserRepository().UserByUUID(ctx, user)
	if err != nil {
		panic(err)
	}

	c.Set("user", user)
}
