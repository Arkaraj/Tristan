package routes

import "github.com/gin-gonic/gin"

func ShowAllUserProjects(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "All Projects",
	})

}

func createProjects(c *gin.Context) {}

func deleteProject(c *gin.Context) {}
