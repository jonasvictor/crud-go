package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasvictor/crud-go/src/configuration/logger"
	"github.com/jonasvictor/crud-go/src/configuration/rest_err/validation"
	"github.com/jonasvictor/crud-go/src/controller/model/request"
	"github.com/jonasvictor/crud-go/src/model"
	"github.com/jonasvictor/crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDoaminInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("jouney", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
