package handlers

import (
	"context"
	"hr-api/cfg"
	"hr-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PsyHandler struct{}

func Psy_fetchAll() gin.HandlerFunc {
	client, err := cfg.ConnectToDB()
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	col := client.Database("substances").Collection("psychwiki")
	cur, err := col.Find(context.Background(), bson.D{{}})
	defer client.Disconnect(context.Background())
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	defer cur.Close(context.Background())
	var result []models.PsyDoc
	for cur.Next(context.Background()) {
		var res models.PsyDoc
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

func Psy_fetchOne(c *gin.Context) {
	var result models.PsyDoc
	drug := c.Param("name")
	client, err := cfg.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	col := client.Database("substances").Collection("psychwiki")
	//opts := options.Find().SetProjection(bson.D{{"_id", 0}})
	err = col.FindOne(context.Background(), bson.D{{"name", drug}}).Decode(&result)
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
