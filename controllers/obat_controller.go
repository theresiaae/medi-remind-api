package controllers

import (
	"medi_remind_backend/config"
	"medi_remind_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateObat - Tambah data obat
func CreateObat(c *gin.Context) {
	var obat models.Obat
	if err := c.ShouldBindJSON(&obat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, obat)
}

// GetAllObat - Ambil semua obat
func GetAllObat(c *gin.Context) {
	var obats []models.Obat
	if err := config.DB.Find(&obats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, obats)
}

// GetObatByID - Ambil 1 obat berdasarkan ID
func GetObatByID(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat

	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, obat)
}

// UpdateObat - Update data obat
func UpdateObat(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat

	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&obat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obat)
}

// DeleteObat - Hapus obat
func DeleteObat(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat

	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Obat berhasil dihapus"})
}
