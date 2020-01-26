package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mahendrahegde/url-shortner-golang/api/models"
	"github.com/mahendrahegde/url-shortner-golang/api/utils"
)

// ShortenUrl godoc
// @Summary shortens the given url
// @Produce json
// @Success 200 {object} models.ShortUrl
// @Router / [post]
// @Param Body body models.ShortUrl true "test"
func (server *Server) ShortenUrl(c *gin.Context) {
	var shortUrl models.ShortUrl
	if err := c.ShouldBindJSON(&shortUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := server.getByKey(shortUrl.Url, shortUrl.GetByUrl)
	if res != (models.ShortUrl{}) {
		c.JSON(http.StatusOK, res)
		return
	}
	if shortUrl.Suffix != "" {
		shortUrl.Short = shortUrl.Suffix
	} else {
		seq, err := server.getNextSeq(shortUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		shortUrl.Short = utils.Base62Encode(seq)
	}
	if err := shortUrl.Insert(server.Db); err != nil {
		log.Println("unable to insert data : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, shortUrl)
}
func (server *Server) GetByUrl(c *gin.Context) {

}
