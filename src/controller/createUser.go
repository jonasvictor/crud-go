package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasvictor/crud-go/src/configuration/logger"
	"github.com/jonasvictor/crud-go/src/configuration/rest_err/validation"
	"github.com/jonasvictor/crud-go/src/controller/model/request"
	"github.com/jonasvictor/crud-go/src/controller/model/response"
	"go.uber.org/zap"
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("jouney", "createUser"))

	c.JSON(http.StatusOK, response)
}
