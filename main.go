package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	URL      string `json:"url"`
	ExpireAt string `json:"expireAt"`
}

type ResponseBody struct {
	ID       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/urls", HandleShortenURL)
	router.GET("/:url", HandleRedirectURL)

	router.Run("localhost:8088")
}

func HandleShortenURL(c *gin.Context) {
	body := RequestBody{}
	if err := c.BindJSON(&body); err != nil {
		return
	}

	expireTime, err := time.Parse(time.RFC3339, body.ExpireAt)
	if err != nil {
		fmt.Println("failed to parse expire time: ", err)
		c.String(400, "invalid expire time format")
		return
	}

	originalURL := body.URL
	url := randSeq(10)

	err = CreateShortenURL(url, originalURL, expireTime)
	if err != nil {
		fmt.Println("failed to insert db: ", err)
		c.String(500, "failed to insert db")
		return
	}

	response := ResponseBody{
		ID:       url,
		ShortUrl: "http://localhost:8088/" + url,
	}

	c.JSON(200, response)
}

func HandleRedirectURL(c *gin.Context) {
	url := c.Param("url")
	originalURL, err := GetOriginalURL(url)
	if err != nil {
		if err == ErrRecordNotFound {
			c.String(404, "not found error")
			return
		} else if err == ErrRecordExpired {
			c.String(404, "record expired")
			return
		}

		fmt.Println("failed to get shorten url: ", err)
		c.String(500, "failed to get db")
		return
	}

	c.Redirect(http.StatusFound, "https://"+originalURL)

	fmt.Println(originalURL)
}
