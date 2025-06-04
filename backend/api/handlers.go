package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"derma-detection/models"
	"derma-detection/services"
)

type Handler struct {
	pythonService *services.PythonService
	uploadDir     string
}

func NewHandler(pythonService *services.PythonService, uploadDir string) *Handler {
	// Upload dizinini oluştur
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		panic(fmt.Sprintf("upload dizini oluşturulamadı: %v", err))
	}

	return &Handler{
		pythonService: pythonService,
		uploadDir:     uploadDir,
	}
}

func (h *Handler) UploadImage(c *gin.Context) {
	// Dosyayı al
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Dosya alınamadı"})
		return
	}

	// Benzersiz dosya adı oluştur
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filepath := filepath.Join(h.uploadDir, filename)

	// Dosyayı kaydet
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Dosya kaydedilemedi"})
		return
	}

	// Python servisini çağır
	prediction, err := h.pythonService.ProcessImage(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: fmt.Sprintf("İşleme hatası: %v", err)})
		return
	}

	c.JSON(http.StatusOK, prediction)
}

func (h *Handler) UploadPDF(c *gin.Context) {
	// Dosyayı al
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Dosya alınamadı"})
		return
	}

	// Benzersiz dosya adı oluştur
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filepath := filepath.Join(h.uploadDir, filename)

	// Dosyayı kaydet
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Dosya kaydedilemedi"})
		return
	}

	// Python servisini çağır
	prediction, err := h.pythonService.ProcessPDF(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: fmt.Sprintf("İşleme hatası: %v", err)})
		return
	}

	c.JSON(http.StatusOK, prediction)
}
