package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BaseURL = "http://localhost:8080" // Update this with your actual domain or localhost configuration
)

func Save(c *gin.Context) {

	playLoad := playLoad{}
	if err := c.ShouldBindJSON(&playLoad); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request - Invalid request payload")
		return
	}

	url, err := saveURL(playLoad.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%s/%s", BaseURL, url.ShortCode))

}

func Open(c *gin.Context) {

	code := c.Param("code")
	url, err := getURL(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.Redirect(http.StatusSeeOther, url.OriginalURL)
}
