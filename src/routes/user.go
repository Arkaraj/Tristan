package routes

import (
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

func ShowUser(c *gin.Context) {

	// filter := bson.D{{"name", "arkaraj"}}

	c.JSON(200, gin.H{
		"message": "Hello User",
	})

}

func AddUserDetails(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello User",
	})
}

func UpdateUserDetails(c *gin.Context) {}

func AddSkillsToUser(c *gin.Context) {}

func AddProjectsToUser(c *gin.Context) {}
