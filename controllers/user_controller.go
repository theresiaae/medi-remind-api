package controllers

import (
	"medi_remind_backend/config"
	"medi_remind_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}
