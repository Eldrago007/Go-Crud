package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/skondrag/go-curd/initializers"
	"github.com/skondrag/go-curd/models"
)

func PostsCreate(c *gin.Context) {
	// Get data of request body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond to them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")

	// Get the Posts
	var post models.Post
	initializers.DB.Find(&post, id)

	// Respond to them

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")

	// Get the data of the req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond to it

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond to it
	c.Status(200)
}
