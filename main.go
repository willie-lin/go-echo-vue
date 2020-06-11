package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"go-echo-vue/database"
	"go-echo-vue/handlers"
	"go-echo-vue/models"
)

func main() {

	db := database.Connect()
	db.AutoMigrate(
		models.Task{},
		models.TaskCollection{},
		)

	_ = db.Close()

	e := echo.New()

	e.File("/","public/index.html")

	e.GET("/ips/:ip", handlers.GetIpInfo())

	e.GET("/tasks", handlers.GetTasks())
	e.GET("/tasks/:id", handlers.FindTaskById())

	e.PUT("/tasks", handlers.PutTasks())
	e.POST("/tasks", handlers.PutTasks())

	e.DELETE("/tasks/:id", handlers.DeleteTasks())

	e.Logger.Fatal(e.Start(":3001"))

}

