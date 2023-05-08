package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type tutorial struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ID          string `json:"id"`
}

var tutorialList = []tutorial{
	{Title: "Tutorial A", Description: "This is tutorial A", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29bc21eb646280e59c84"},
	{Title: "Tutorial B", Description: "This is tutorial B", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29c2acefe3e62a866731"},
	{Title: "Tutorial C", Description: "This is tutorial C", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f2a0eacefe3e62a866735"},
}

func main() {

	// /*
	//         List databases
	// */
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println(databases)

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
