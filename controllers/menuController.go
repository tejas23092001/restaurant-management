package controller

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := menuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing the menu items"})
			return
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the menu"})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menu.Created_at = time.Now()
		menu.Updated_at = time.Now()
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu item was not created"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.Before(check) && end.After(start)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start or end date"})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})
		}

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}

		menu.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

		upsert := true
		opt := options.UpdateOptions{Upsert: &upsert}

		result, err := menuCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu update failed"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
