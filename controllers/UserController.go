package controllers

import (
	"GoBlog/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func UsersLogin(c *gin.Context) {
	user := data.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.IsAuthenticated(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "authantication error",
	})
}
func UsersRegister(c *gin.Context){
 user := 	data.User{}
 err := c.ShouldBindJSON(&user)
 if err != nil{
 	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	 return
 }
 db , _ :=c.Get("db")
 conn := db.(pgx.Conn)
 err = user.Register(&conn)

 if err != nil{
 	fmt.Println("error")
 	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	 return
 }
	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
	})
}