package main

import (
	"GoBlog/controllers"
	"github.com/gin-gonic/gin"
)

func routing() {
	conn, err := connectToDb()
	if err != nil {
		return
	}
	router := gin.Default()
	router.Use(databaseMiddleware(*conn))
	userEndpoints := router.Group("users")
	{
		userEndpoints.POST("register", controllers.UsersRegister)
		userEndpoints.POST("login", controllers.UsersLogin)
	}
	postEndpoints := router.Group("posts")
	{
		postEndpoints.GET("index", controllers.PostsIndex)
		postEndpoints.GET("index/:id", controllers.PostByID)
		postEndpoints.POST("create", authenticationMiddleware(), controllers.PostsCreate)
		postEndpoints.GET("myposts", authenticationMiddleware(), controllers.PostsByCurrentUser)
		postEndpoints.PUT("update", authenticationMiddleware(), controllers.PostsUpdate)
		postEndpoints.DELETE("delete/:id", authenticationMiddleware(), controllers.PostsDelete)
	}
	err = router.Run(":52")
}
