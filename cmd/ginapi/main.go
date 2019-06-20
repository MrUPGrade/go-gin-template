package main

import (
	dbConfig "ginapi/pkg/db/config"
	models "ginapi/pkg/db/models"
	"ginapi/pkg/resources"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func registerRoutes(db *gorm.DB, engine *gin.Engine) {
	engine.GET("/ping", ping)

	rb := resources.ResourceBase{DB: db}

	todoResource := resources.TodoController{ResourceBase: rb}
	engine.GET("/todo", todoResource.Get)
}

func main() {
	db, err := dbConfig.ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_ = db.LogMode(true)
	db.SingularTable(true)
	_ = db.AutoMigrate(&models.Todo{})

	r := gin.Default()

	registerRoutes(db, r)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
