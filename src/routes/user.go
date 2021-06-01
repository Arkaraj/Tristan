package routes

import (
	"Tristan/src/database"
	"Tristan/src/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")

func ShowUser(c *gin.Context) {

	var me models.User

	filter := bson.M{"name": bson.M{"$eq": "Arkaraj"}}

	lookupQuery1 := bson.D{{"$lookup", bson.D{{"from", "Skill"}, {"localField", "skills"}, {"foreignField", "_id"}, {"as", "mySkills"}}}}
	lookupQuery2 := bson.D{{"$lookup", bson.D{{"from", "Project"}, {"localField", "projects"}, {"foreignField", "_id"}, {"as", "myProjects"}}}}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&me)

	showLoadedCursor, err := userCollection.Aggregate(context.TODO(), mongo.Pipeline{lookupQuery1, lookupQuery2})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	} else {

		var userLoaded []bson.M

		if err = showLoadedCursor.All(context.TODO(), &userLoaded); err != nil {
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"my Profile": userLoaded,
			})
		}
	}

}

func AddUserDetails(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	user.Id = primitive.NewObjectID()

	user.Skills = []primitive.ObjectID{}
	user.Projects = []primitive.ObjectID{}

	result, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		defer cancel()
		return
	}
	defer cancel()

	//return the id of the created object to the frontend
	c.JSON(http.StatusOK, gin.H{
		"User":       user,
		"InsertedId": result.InsertedID,
	})
}

func UpdateUserDetails(c *gin.Context) {}

func AddSkillsToUser(c *gin.Context) {
	paramUserId := c.Param("userId")
	paramSkillId := c.Param("skillId")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(paramUserId)
	skillId, err := primitive.ObjectIDFromHex(paramSkillId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		var me models.User

		err := userCollection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": userId}}).Decode(&me)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		skill := me.Skills

		skill = append(skill, skillId)

		filter := bson.M{"_id": bson.M{"$eq": userId}}
		update := bson.M{"$set": bson.M{"skills": skill}}

		// var result bson.M
		user, err := userCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			me.Skills = append(me.Skills, skillId)
			c.JSON(http.StatusOK, gin.H{
				"update":     user,
				"My Profile": me,
			})
		}

	}
}

func AddProjectsToUser(c *gin.Context) {
	paramUserId := c.Param("userId")
	paramProjectId := c.Param("projectId")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(paramUserId)
	projectId, err := primitive.ObjectIDFromHex(paramProjectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		var me models.User

		err := userCollection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": userId}}).Decode(&me)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		project := me.Projects

		project = append(project, projectId)

		filter := bson.M{"_id": bson.M{"$eq": userId}}
		update := bson.M{"$set": bson.M{"projects": project}}

		// var result bson.M
		userUpdate, err := userCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			me.Projects = append(me.Projects, projectId)
			c.JSON(http.StatusOK, gin.H{
				"update":     userUpdate,
				"My Profile": me,
			})
		}

	}
}
