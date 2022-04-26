package middlewares

import (
	"my-gram/config"
	"my-gram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.InitDB()
		getId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Photo := models.Photo{}

		if err := db.Preload("User").Preload("Comments").First(&Photo, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "data not found",
				"message": err.Error(),
			})
			return
		}

		if uint(UserId) != Photo.User_id {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you can't access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.InitDB()
		getId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Comment := models.Comment{}

		if err := db.Preload("User").Preload("Photo").First(&Comment, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "data not found",
				"message": err.Error(),
			})
			return
		}

		if uint(UserId) != Comment.User_id {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you can't access this data",
			})
			return
		}
		c.Next()
	}
}

func MediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.InitDB()
		getId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		Media := models.Media{}

		if err := db.Preload("User").First(&Media, getId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "data not found",
				"message": err.Error(),
			})
			return
		}

		if int(UserId) != int(Media.User_id) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you can't access this data",
			})
			return
		}
		c.Next()
	}
}
