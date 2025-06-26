package analyzer

import (
	"errors"
	"fmt"
)

type FileNotFoundError struct {
	FilePath string
	Err      error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.FilePath)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

type ParseError struct {
	LogID   string
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing pour le log '%s': %s", e.LogID, e.Message)
}

func NewFileNotFoundError(filePath string, originalErr error) *FileNotFoundError {
	return &FileNotFoundError{
		FilePath: filePath,
		Err:      originalErr,
	}
}

func NewParseError(logID, message string) *ParseError {
	return &ParseError{
		LogID:   logID,
		Message: message,
	}
}

func IsFileNotFoundError(err error) bool {
	var fileErr *FileNotFoundError
	return errors.As(err, &fileErr)
}

func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}
