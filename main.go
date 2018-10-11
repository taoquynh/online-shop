package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/online-shop/config"
	"github.com/online-shop/controller"
	"github.com/online-shop/router"
	"github.com/online-shop/model"
)

func main() {
	r := gin.Default()

	// Đọc cấu hình từ file json
	config := config.SetupConfig()


	// Khởi tạo controller
	c := controller.NewController()
	c.Config = config


	// Kết nối CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db
	log.Println(config.Database.Database)
	err := model.MigrationDb(db, config.ServiceName)
	if err != nil {
		panic(err)
	}

	router.SetupRouter(config, r, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
