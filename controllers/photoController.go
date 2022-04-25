package controllers

import (
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoRepo struct {
	DB *gorm.DB
}

func (p *PhotoRepo) GetPhoto(c *gin.Context) {
	Photos := []models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "data not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})

}
func (p *PhotoRepo) UploadPhoto(c *gin.Context) {
	Photo := models.Photo{}
	contextType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_id = uint(userId)
	Photo.Created_at = time.Now()
	Photo.Updated_at = time.Now()

	if err := p.DB.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to updload photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.Id,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.User_id,
		"created_at": Photo.Created_at,
	})
}
