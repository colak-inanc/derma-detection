package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"derma-detection/models"
)

type PythonService struct {
	baseURL string
}

func NewPythonService(baseURL string) *PythonService {
	return &PythonService{
		baseURL: baseURL,
	}
}

func (s *PythonService) ProcessImage(imagePath string) (*models.PredictionResponse, error) {
	// Dosyayı aç
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("dosya açılamadı: %v", err)
	}
	defer file.Close()

	// Multipart form oluştur
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Dosyayı forma ekle
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		return nil, fmt.Errorf("form oluşturulamadı: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("dosya kopyalanamadı: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("form kapatılamadı: %v", err)
	}

	// HTTP isteği gönder
	req, err := http.NewRequest("POST", s.baseURL+"/process-image", body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturulamadı: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("istek gönderilemedi: %v", err)
	}
	defer resp.Body.Close()

	// Yanıtı işle
	var prediction models.PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return nil, fmt.Errorf("yanıt işlenemedi: %v", err)
	}

	return &prediction, nil
}

func (s *PythonService) ProcessPDF(pdfPath string) (*models.PredictionResponse, error) {
	// PDF işleme mantığı burada implement edilecek
	// Şimdilik örnek bir yanıt döndürüyoruz
	return &models.PredictionResponse{
		DiseaseProbability: 0.85,
		PredictedDisease:   "Akne",
		ConfidenceScore:    0.92,
		RecommendedActions: []string{
			"Dermatoloğa başvurun",
			"Reçetesiz akne kremi kullanın",
			"Yüzünüzü günde iki kez yıkayın",
		},
	}, nil
}
