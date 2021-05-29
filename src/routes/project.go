package routes

import (
	"github.com/gin-gonic/gin"
)

func ShowAllUserProjects(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "All Projects",
	})

}

func CreateProjects(c *gin.Context) {}

func DeleteProject(c *gin.Context) {}
