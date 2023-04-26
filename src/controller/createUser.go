package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasvictor/crud-go/src/configuration/logger"
	"github.com/jonasvictor/crud-go/src/configuration/rest_err/validation"
	"github.com/jonasvictor/crud-go/src/controller/model/request"
	"github.com/jonasvictor/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDoaminInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("jouney", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("jouney", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("jouney", "createUser"))

	c.String(http.StatusOK, "")
}
