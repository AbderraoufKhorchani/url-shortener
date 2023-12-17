package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BaseURL = "http://localhost:8080" // Update this with your actual domain or localhost configuration
)

// Save godoc
// @Summary Save a URL and generate a shortcode
// @Description Save a URL and generate a shortcode
// @Tags URL Shortener
// @Accept json
// @Produce json
// @Param request body playLoad true "Request payload containing the URL"
// @Success 200 {string} string "Returns the shortened URL"
// @Failure 400 {string} string "Bad Request - Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /generate [post]
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

// Open godoc
// @Summary Get the original URL based on the provided shortcode
// @Description Get the original URL based on the provided shortcode
// @Tags URL Shortener
// @Accept json
// @Produce json
// @Param code path string true "Shortcode to be used for redirection"
// @Success 200 {object} RedirectResponse "Returns the original URL"
// @Failure 500 {string} string "Internal Server Error"
// @Router /{code} [get]
func Open(c *gin.Context) {
	code := c.Param("code")
	url, err := getURL(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Check if the request is coming from Swagger UI
	if c.GetHeader("Accept") == "application/json" {
		// Return the original URL in the response
		c.JSON(http.StatusOK, url.OriginalURL)
	} else {
		// Perform the actual redirection
		c.Redirect(http.StatusSeeOther, url.OriginalURL)
	}
}
