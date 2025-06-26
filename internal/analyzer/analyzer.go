package analyzer

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/phil51297/log_analyzer/internal/config"
)

type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLog(logConfig config.LogConfig) AnalysisResult {
	result := AnalysisResult{
		LogID:        logConfig.ID,
		FilePath:     logConfig.Path,
		Status:       "OK",
		Message:      "Analyse terminée avec succès.",
		ErrorDetails: "",
	}

	if err := checkFileAccess(logConfig.Path); err != nil {
		result.Status = "FAILED"
		result.ErrorDetails = err.Error()

		if IsFileNotFoundError(err) {
			result.Message = "Fichier introuvable."
		} else {
			result.Message = "Erreur d'accès au fichier."
		}
		return result
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	analysisDelay := time.Duration(50+r.Intn(150)) * time.Millisecond
	time.Sleep(analysisDelay)

	if r.Float32() < 0.1 {
		parseErr := NewParseError(logConfig.ID, "format de log invalide détecté")
		result.Status = "FAILED"
		result.Message = "Erreur lors du parsing du log."
		result.ErrorDetails = parseErr.Error()
		return result
	}

	return result
}

func checkFileAccess(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return NewFileNotFoundError(filePath, err)
		}
		return NewFileNotFoundError(filePath, err)
	}

	if info.IsDir() {
		return NewFileNotFoundError(filePath, errors.New("le chemin pointe vers un répertoire"))
	}

	file, err := os.Open(filePath)
	if err != nil {
		return NewFileNotFoundError(filePath, err)
	}
	defer file.Close()

	return nil
}
