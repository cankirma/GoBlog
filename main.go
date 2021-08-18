package main

import (
	"GoBlog/data"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)



func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {

	routing()
}

func connectToDb() (c *pgx.Conn, err error) {
	conStr, _ := os.LookupEnv("CONNECTION_STRING")
	conn, err := pgx.Connect(context.Background(), conStr)
	if err != nil {
		fmt.Println("unable connect to db")
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func databaseMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(401, gin.H{"errror": "not authenticated"})
			c.Abort()
			return
		}
		token := split[1]
		isValid, userId := data.IsTokenValid(token)
		if isValid == false {
			c.JSON(401, gin.H{"error": "not authenticated"})
			c.Abort()
		} else {
			c.Set("user_id", userId)
			c.Next()
		}

	}
}

