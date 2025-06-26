package reporter

import (
	"encoding/json"
	"os"

	"github.com/phil51297/log_analyzer/internal/analyzer"
)

func ExportToJSON(results []analyzer.AnalysisResult, filePath string) error {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func ExportToJSONString(results []analyzer.AnalysisResult) (string, error) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
