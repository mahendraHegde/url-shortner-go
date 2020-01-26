package api

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mahendrahegde/url-shortner-golang/api/models"
	"github.com/mahendrahegde/url-shortner-golang/api/utils"
)

func (server *Server) getNextSeq(shortUrl models.ShortUrl) (int, error) {
	seq, _ := server.Cache.Get("sequence").Result()
	var nextSeq int
	//if sequence doesn't exist initialize it with current DB state
	if seq == "" {
		res, _ := shortUrl.GetLastUrlSeq(server.Db)
		if res.Short != "" {
			lastSeq, err := utils.Base62Decode(res.Short)

			if err != nil {
				startSeq, err := strconv.Atoi(server.ENV.START_SEQ)
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
			startSeq, err := strconv.Atoi(server.ENV.START_SEQ)
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

func (server *Server) getByKey(url string, f func(*gorm.DB, string) (models.ShortUrl, error)) models.ShortUrl {
	val, _ := server.Cache.Get(url).Result()
	var res models.ShortUrl
	if val != "" {
		if err := json.Unmarshal([]byte(val), &res); err == nil {
			return res
		}
	}
	if res, err := f(server.Db, strings.Split(url, "-")[1]); err == nil {
		if marshalled, err := json.Marshal(res); err == nil {
			server.Cache.Set(url, marshalled, 0)
		}
		return res
	}
	return res
}
