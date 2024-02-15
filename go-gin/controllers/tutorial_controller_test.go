package controllers

import (
	"encoding/json"
	"go-gin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestSearchTutorials(t *testing.T) {

	t.Run("Search by title", func(t *testing.T) {
		// Arrange
		title := "asdfasdf"
		req, _ := http.NewRequest("GET", "/tutorials?title="+title, nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = req

		// Act
		SearchTutorials()
		res := w.Result()
		defer res.Body.Close()

		// Assert
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var tutorials []models.Tutorial
		json.NewDecoder(res.Body).Decode(&tutorials)

		assert.Greater(t, len(tutorials), 0)
		assert.Equal(t, title, tutorials[0].Title)
	})

	t.Run("Search no results", func(t *testing.T) {
		// Arrange
		title := "Fake Tutorial"
		req, _ := http.NewRequest("GET", "/tutorials?title="+title, nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = req

		// Act
		SearchTutorials()
		res := w.Result()
		defer res.Body.Close()

		// Assert
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var tutorials []models.Tutorial
		json.NewDecoder(res.Body).Decode(&tutorials)

		assert.Equal(t, 0, len(tutorials))
	})
}
