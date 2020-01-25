package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortenUrlBody struct {
	Url string `json:"url" binding:"required"`
}
// ShortenUrl godoc
// @Summary shortens the given url
// @Produce json
// @Success 200 {object} ShortenUrlBody
// @Router / [post]
// @Param Body body ShortenUrlBody true "test"
func ShortenUrl(c *gin.Context) {
	var body ShortenUrlBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, body)
}
