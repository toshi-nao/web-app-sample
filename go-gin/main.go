package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)


// album represents data about a record album.
type tutorial struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ID          string `json:"id"`
}

// albums slice to seed record album data.
var tutorialList = []tutorial{
	{Title: "Tutorial 1", Description: "This is tutorial 1", Published: false, 
    CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29bc21eb646280e59c84"},
	{Title: "Tutorial 2", Description: "This is tutorial 2", Published: false, 
    CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29c2acefe3e62a866731"},
	{Title: "Tutorial 3", Description: "This is tutorial 3", Published: false, 
    CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f2a0eacefe3e62a866735"},
}

func main() {
	router := gin.Default()
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{
        "http://localhost:8081",
    }
    router.Use(cors.New(config))
	router.GET("/api/tutorials", getAllTutorials)
	router.Run("localhost:8080")
}


func getAllTutorials(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tutorialList)
    println("getAllTutorials")
}
