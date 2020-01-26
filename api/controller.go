package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mahendrahegde/url-shortner-golang/api/models"
	"github.com/mahendrahegde/url-shortner-golang/api/utils"
)

const START_SEQ = "START_SEQ"

func (server *Server) getNextSeq(shortUrl models.ShortUrl) (int, error) {
	seq, _ := server.Cache.Get("sequence").Result()
	var nextSeq int
	//if sequence doesn't exist initialize it with current DB state
	if seq == "" {
		res, _ := shortUrl.GetLastUrlSeq(server.Db)
		if res.Short != "" {
			lastSeq, err := utils.Base62Decode(res.Short)

			if err != nil {
				startSeq, err := strconv.Atoi(os.Getenv(START_SEQ))
				if err != nil {
					log.Fatal("unable to parse start seq ", err.Error())
					return -1, err
				} else {
					nextSeq = startSeq
				}
			} else {
				nextSeq = int(lastSeq + 1)
			}
		} else {
			startSeq, err := strconv.Atoi(os.Getenv(START_SEQ))
			if err != nil {
				log.Fatal("unable to parse start seq ", err.Error())
				return -1, err
			} else {
				nextSeq = startSeq
			}
		}

	} else {
		nextSeq, err := strconv.Atoi(seq)
		if err != nil {
			log.Fatal("unable to parse start seq ", err.Error())
			return -1, err
		}
		if err := server.Cache.Set("sequence", nextSeq+1, 0).Err(); err != nil {
			log.Println("unable to set sequence ", nextSeq)
		}
		return nextSeq, nil
	}
	if err := server.Cache.Set("sequence", nextSeq+1, 0).Err(); err != nil {
		log.Println("unable to set sequence ", nextSeq)
	}
	return nextSeq, nil
}

// ShortenUrl godoc
// @Summary shortens the given url
// @Produce json
// @Success 200 {object} models.ShortUrl
// @Router / [post]
// @Param Body body models.ShortUrl true "test"
func (server *Server) ShortenUrlController(c *gin.Context) {
	var shortUrl models.ShortUrl
	if err := c.ShouldBindJSON(&shortUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
