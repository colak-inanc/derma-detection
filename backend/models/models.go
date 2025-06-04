package models

type PredictionRequest struct {
	ImagePath string   `json:"image_path"`
	CBCData   CBCData  `json:"cbc_data"`
	Symptoms  []string `json:"symptoms"`
}

type CBCData struct {
	HCT float64 `json:"HCT"`
	MCV float64 `json:"MCV"`
	PDW float64 `json:"PDW"`
}

type PredictionResponse struct {
	DiseaseProbability float64  `json:"disease_probability"`
	PredictedDisease   string   `json:"predicted_disease"`
	ConfidenceScore    float64  `json:"confidence_score"`
	RecommendedActions []string `json:"recommended_actions"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
