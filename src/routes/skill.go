package routes

import "github.com/gin-gonic/gin"

func ShowAllUserSkills(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "All Skills",
	})

}

func createSkills(c *gin.Context) {}

func deleteSkill(c *gin.Context) {}
