package main

import (
	"fmt"
	"main/api"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://mos2022.herokuapp.com/"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()

	api.Setup(router)

	// router.Run(":8081")

	// In case of running on Heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("Running on Heroku using random PORT")
		router.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		router.Run(":" + port)
	}
}
