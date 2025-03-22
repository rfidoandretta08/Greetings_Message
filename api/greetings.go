package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Struktur Greeting
type Greeting struct {
	ID      int    `json:"id"`
	Message string `json:"message" binding:"required"`
	Status  string `json:"status"`
}

// Data awal
var greetings = []Greeting{
	{ID: 1, Message: "Selamat pagi", Status: "active"},
}

// Inisialisasi router
func initRouter() *gin.Engine {
	router := gin.Default()

	Salam := router.Group("/api/Salam")
	{
		Salam.GET("/greetings", getGreetings)
		Salam.GET("/greetings/:id", getGreetingByID)
		Salam.POST("/greetings", createGreeting)
	}

	return router
}

// Handler utama untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := initRouter()
	router.ServeHTTP(w, r)
}

// GET /api/Salam/greetings
func getGreetings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": greetings})
}

// GET /api/Salam/greetings/:id
func getGreetingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, g := range greetings {
		if g.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": g})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Greeting not found"})
}

// POST /api/Salam/greetings
func createGreeting(c *gin.Context) {
	var newGreeting Greeting

	if err := c.ShouldBindJSON(&newGreeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGreeting.ID = len(greetings) + 1
	newGreeting.Status = "active"

	greetings = append(greetings, newGreeting)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Greeting created successfully",
		"data":    newGreeting,
	})
}
