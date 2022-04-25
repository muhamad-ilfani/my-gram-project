package middlewares

import (
	"my-gram/config"
	"my-gram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.InitDB()
		getId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
		}
		UserData := c.MustGet("userData").(jwt.MapClaims)
		UserId := UserData["id"].(float64)
		User := models.User{}

		if err := db.Preload("Photos").Preload("Comments").Preload("Medias").Where("id=?", getId).Take(&User).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "data not found",
				"message": err.Error(),
			})
			return
		}

		if int(UserId) != User.Id {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you can't access this data",
			})
			return
		}
		c.Next()
	}
}
