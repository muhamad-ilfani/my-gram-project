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

type CommentRepo struct {
	DB *gorm.DB
}

func (o *CommentRepo) GetComment(c *gin.Context) {
	Comments := []models.Comment{}

	if err := o.DB.Debug().Preload("User").Preload("Photo").Find(&Comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}
func (o *CommentRepo) UploadComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	if err := o.DB.Debug().Find(&models.Comment{}, Comment.Photo_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}
	Comment.User_id = uint(userId)
	Comment.Created_at = time.Now()
	Comment.Updated_at = time.Now()

	if err := o.DB.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":   "failed to upload comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Comment,
	})
}

func (o *CommentRepo) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	OldComment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.User_id = uint(userId)
	Comment.Updated_at = time.Now()

	if err := o.DB.Debug().First(&OldComment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Comment not found",
			"messsage": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Model(&OldComment).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         OldComment.Id,
		"message":    OldComment.Message,
		"user_id":    OldComment.User_id,
		"updated_at": OldComment.Updated_at,
	})
}

func (o *CommentRepo) DeleteComment(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	if err := o.DB.Debug().First(&Comment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Delete(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})

}
