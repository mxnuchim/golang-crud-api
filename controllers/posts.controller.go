package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mxnuchim/golang-crud-api/initializers"
	"github.com/mxnuchim/golang-crud-api/models"
)

func CreatePost(c *gin.Context) {
	// get data from request body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	//handle and log error
	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	//return created post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	//get the posts
	var posts []models.Post

	initializers.DB.Find(&posts)

	//respond with the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	//get post id off the url
	id := c.Param("id")

	//get the post
	var post models.Post

	initializers.DB.First(&post, id)

	//respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// get post id off the url
	id := c.Param("id")	

	// get data from request body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	// find the post to update
	var post models.Post

	initializers.DB.First(&post, id)

	// update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// get post id off the url
	id := c.Param("id")	

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
