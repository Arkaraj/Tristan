package routes

import "github.com/gin-gonic/gin"

func ShowUser(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello User",
	})

}

func addUserDetails(c *gin.Context) {}

func updateUserDetails(c *gin.Context) {}

func addSkillsToUser(c *gin.Context) {}

func addProjectsToUser(c *gin.Context) {}
