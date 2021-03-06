package routes

import (
	"Tristan/src/database"
	"Tristan/src/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection = database.OpenCollection(database.Client, "Project")

func ShowAllUserProjects(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// filter := bson.M{}

	lookupQuery := bson.D{{"$lookup", bson.D{{"from", "Skill"}, {"localField", "skills"}, {"foreignField", "_id"}, {"as", "projectSkills"}}}}

	cursor, err := projectCollection.Aggregate(context.TODO(), mongo.Pipeline{lookupQuery})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cursor.Close(ctx)
	} else {
		var projectLoaded []bson.M

		if err = cursor.All(context.TODO(), &projectLoaded); err != nil {
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"Projects": projectLoaded,
			})
		}

		/* for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				defer cursor.Close(ctx)
			}
			c.JSON(http.StatusOK, result)
		} */

	}

}

func CreateProjects(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var project models.Project

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	project.Id = primitive.NewObjectID()

	project.Skills = []primitive.ObjectID{}

	project.StartTime.Format("2006-January-02")
	project.EndTime.Format("2006-January-02")

	result, insertErr := projectCollection.InsertOne(ctx, project)
	if insertErr != nil {
		msg := fmt.Sprintf("The Project was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		defer cancel()
		return
	}
	defer cancel()

	//return the id of the created object to the frontend
	c.JSON(http.StatusOK, gin.H{
		"message":    "New Project Created",
		"Project":    project,
		"InsertedId": result.InsertedID,
	})

}

func AddSkillsToProjects(c *gin.Context) {

	paramProjectId := c.Param("projectId")
	paramSkillId := c.Param("skillId")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projectId, err := primitive.ObjectIDFromHex(paramProjectId)
	skillId, err := primitive.ObjectIDFromHex(paramSkillId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		var work models.Project

		err := projectCollection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": projectId}}).Decode(&work)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		skill := work.Skills

		skill = append(skill, skillId)

		filter := bson.M{"_id": bson.M{"$eq": projectId}}
		update := bson.M{"$set": bson.M{"skills": skill}}

		// var result bson.M
		project, err := projectCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			work.Skills = append(work.Skills, skillId)
			c.JSON(http.StatusOK, gin.H{
				"update":  project,
				"project": work,
			})
		}

	}

}

func DeleteProject(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Param("projectId")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		filter := bson.M{"_id": idPrimitive}

		result, err := projectCollection.DeleteOne(ctx, filter)

		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		} else {
			if result.DeletedCount == 0 {
				c.JSON(http.StatusOK, gin.H{
					"msg": "No _id found with to delete",
				})
			} else {

				c.JSON(http.StatusOK, gin.H{
					"result": result,
				})

			}
		}

	}

}
