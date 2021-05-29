package main

import (
	"context"
	"fmt"
	"goRestAPI/src/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is a simple Portfolio REST API",
	})
}

func main() {
	server := gin.Default()

	// From mongo-go-driver docs
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/Portfolio"))

	// erro := client.Ping(ctx, readpref.Primary())

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	server.GET("/api/", api)

	// Groups
	user := server.Group("/api/user")
	{
		user.GET("/", routes.ShowUser)
	}

	skill := server.Group("/api/skill")
	{
		skill.GET("/", routes.ShowAllUserSkills)
	}

	projects := server.Group("/api/project")
	{
		projects.GET("/", routes.ShowAllUserProjects)
	}

	port := ":7000"

	server.Run(port) // Runs on port http://localhost:8080 (by default)
}
