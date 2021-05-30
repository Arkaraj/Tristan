package routes

import (
	"context"
	"fmt"
	"goRestAPI/src/database"
	"goRestAPI/src/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var skillCollection *mongo.Collection = database.OpenCollection(database.Client, "Skill")

func ShowAllUserSkills(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := skillCollection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cursor.Close(ctx)
	}
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cursor.Close(ctx)
		}
		c.JSON(http.StatusOK, result)
	}

}

func CreateSkills(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	var skill models.Skill

	// bind the object that comes in with the declared varaible. thrrow an error if one occurs
	if err := c.BindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	//generate new ID for the object to be created
	skill.Id = primitive.NewObjectID()

	//insert the newly created object into mongodb
	result, insertErr := skillCollection.InsertOne(ctx, skill)
	if insertErr != nil {
		msg := fmt.Sprintf("The Skill was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		defer cancel()
		return
	}
	defer cancel()

	//return the id of the created object to the frontend
	c.JSON(http.StatusOK, gin.H{
		"message":    "New Skill Created",
		"Skill":      skill,
		"InsertedId": result.InsertedID,
	})

}

func DeleteSkill(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Param("skillId")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		filter := bson.M{"_id": idPrimitive}

		result, err := skillCollection.DeleteOne(ctx, filter)

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
