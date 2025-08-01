package handlers

import (
	"context"
	"hr-api/cfg"
	"hr-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SelfHandler struct{}

func Self_updateOne(c *gin.Context) {
	var updateBody models.SelfModel
	if err := c.BindJSON(&updateBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	drug := c.Param("name")
	client, err := cfg.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	col := client.Database("substances").Collection("substances")
	result, err := col.UpdateOne(context.Background(), bson.D{{"name", drug}}, bson.D{{"$set", updateBody}})
	defer client.Disconnect(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Self_deleteOne(c *gin.Context) {
	//var result models.SelfModel
	drug := c.Param("name")
	client, err := cfg.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	col := client.Database("substances").Collection("substances")
	result, err := col.DeleteOne(context.Background(), bson.D{{"name", drug}})
	defer client.Disconnect(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Self_insertOne(c *gin.Context) {
	var bodydata models.SelfModel
	if err := c.BindJSON(&bodydata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := cfg.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	col := client.Database("substances").Collection("substances")
	result, err := col.InsertOne(context.Background(), bodydata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Self_fetchAll() gin.HandlerFunc {
	client, err := cfg.ConnectToDB()
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	col := client.Database("substances").Collection("substances")
	opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	cur, err := col.Find(context.Background(), bson.D{{}}, opts)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	defer client.Disconnect(context.Background())
	defer cur.Close(context.Background())
	var result []models.SelfModel
	for cur.Next(context.Background()) {
		var res models.SelfModel
		if err := cur.Decode(&res); err != nil {
			return func(c *gin.Context) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
		result = append(result, res)
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, result)
	}
}

func Self_fetchOne(c *gin.Context) {
	var result models.SelfModel
	drug := c.Param("name")
	client, err := cfg.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bad request": err.Error()})
	}
	col := client.Database("substances").Collection("substances")
	filter := bson.D{{"name", bson.D{{"$regex", "^" + drug + "$"}, {"$options", "i"}}}}
	err = col.FindOne(context.Background(), filter).Decode(&result)
	//err = cur.Decode(&result)
	defer client.Disconnect(context.Background())
	switch err {
	case mongo.ErrNoDocuments:
		c.JSON(http.StatusNotFound, gin.H{"Drug not found": err.Error()})
		return
	case mongo.ErrClientDisconnected:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	default:
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
		return
	}
}
