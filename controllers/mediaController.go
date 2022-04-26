package controllers

import (
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MediaRepo struct {
	DB *gorm.DB
}

func (m *MediaRepo) UploadMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Media := models.Media{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.User_id = uint(userId)
	Media.Created_at = time.Now()
	Media.Updated_at = time.Now()

	if err := m.DB.Debug().Create(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to upload social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Media.Id,
		"name":             Media.Name,
		"social_media_url": Media.Social_media_url,
		"user_id":          Media.User_id,
		"created_at":       Media.Created_at,
	})
}

func (m *MediaRepo) GetMedia(c *gin.Context) {
	Medias := []models.Media{}

	if err := m.DB.Debug().Find(&Medias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "can't find media",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"media": Medias,
	})
}

func (m *MediaRepo) UpdateMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("socialMediaId"))

	Media := models.Media{}
	OldMedia := models.Media{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Media)
	} else {
		c.ShouldBind(&Media)
	}

	Media.Updated_at = time.Now()
	Media.User_id = uint(userId)

	if err := m.DB.Debug().First(&OldMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Model(&OldMedia).Updates(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               OldMedia.Id,
		"name":             OldMedia.Name,
		"social_media_url": OldMedia.Social_media_url,
		"user_id":          OldMedia.User_id,
		"updated_at":       OldMedia.Updated_at,
	})

}

func (m *MediaRepo) DeleteMedia(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("socialMediaId"))
	Media := models.Media{}

	if err := m.DB.Debug().First(&Media, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Delete(&Media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete media",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
