package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jonasvictor/crud-go/src/configuration/logger"
	"github.com/jonasvictor/crud-go/src/controller"
	"github.com/jonasvictor/crud-go/src/controller/routes"
	"github.com/jonasvictor/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start the application")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
