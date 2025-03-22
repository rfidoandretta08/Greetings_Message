package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Greeting struct {
	ID      int    `json:"id"`
	Message string `json:"message" binding:"required"`
	Status  string `json:"status"`
}

var greetings = []Greeting{
	{ID: 1, Message: "Selamat pagi", Status: "active"},
}

func Handler() *gin.Engine {
	router := gin.Default()

	Salam := router.Group("/api/Salam")
	{
		Salam.GET("/greetings", getGreetings)
		Salam.GET("/greetings/:id", getGreetingByID)
		Salam.POST("/greetings", createGreeting)
	}

	return router
}

func getGreetings(c *gin.Context) {
	c.JSON(200, gin.H{"data": greetings})
}

func getGreetingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, g := range greetings {
		if g.ID == id {
			c.JSON(200, gin.H{"data": g})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Greeting not found"})
}

func createGreeting(c *gin.Context) {
	var newGreeting Greeting

	if err := c.ShouldBindJSON(&newGreeting); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newGreeting.ID = len(greetings) + 1
	newGreeting.Status = "active"

	greetings = append(greetings, newGreeting)

	c.JSON(201, gin.H{
		"message": "Greeting created successfully",
		"data":    newGreeting,
	})
}
