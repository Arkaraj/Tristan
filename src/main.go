package main

import (
	"goRestAPI/src/routes"

	"github.com/gin-gonic/gin"
)

func api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is a simple Portfolio REST API",
	})
}

func main() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	server.GET("/api/", api)

	// Groups
	user := server.Group("/api/me")
	{
		user.GET("/", routes.ShowUser)
		user.POST("/", routes.AddUserDetails)
		user.PUT("/", routes.UpdateUserDetails)
		user.POST("/skill/:userId/:skillId/", routes.AddSkillsToUser)
		user.POST("/project/:userId/:projectId", routes.AddProjectsToUser)
	}

	skill := server.Group("/api/skill")
	{
		skill.GET("/", routes.ShowAllUserSkills)

		skill.POST("/", routes.CreateSkills)

		skill.DELETE("/:skillId", routes.DeleteSkill)
	}

	projects := server.Group("/api/project")
	{
		projects.GET("/", routes.ShowAllUserProjects)
		projects.POST("/", routes.CreateProjects)
		projects.PUT("/:projectId/:skillId", routes.AddSkillsToProjects)
		projects.DELETE("/:projectId", routes.DeleteProject)
	}

	port := ":7000"

	// Runs on port http://localhost:8080 by default
	server.Run(port)
}
