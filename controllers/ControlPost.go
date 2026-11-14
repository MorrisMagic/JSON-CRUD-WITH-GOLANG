package controllers

import (
	"crud/initializers"
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type notePost struct {
	Body  string
	Title string
}

func Posts(c *gin.Context) {
	var request notePost

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't read data",
		})
	}

	post := models.Post{Title: request.Title, Body: request.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(201, gin.H{
		"posts": posts,
	})

}
func OnePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, id)

	c.JSON(201, gin.H{
		"post": post,
	})

}
func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var todo notePost

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't read data",
		})
	}
	var post models.Post

	initializers.DB.Find(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: todo.Title,
		Body:  todo.Body,
	})

	c.JSON(201, gin.H{
		"post": post,
	})

}
func DeletePost(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}
