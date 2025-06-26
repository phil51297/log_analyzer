// Package analyzer fournit les types d'erreurs personnalisées pour l'analyse de logs.
// Ces erreurs permettent une gestion fine des cas d'échec avec des messages localisés.
package analyzer

import (
	"errors"
	"fmt"
)

// FileNotFoundError représente une erreur de fichier introuvable ou inaccessible.
// Elle implémente l'interface error et supporte le wrapping d'erreurs.
type FileNotFoundError struct {
	FilePath string // Chemin du fichier qui pose problème
	Err      error  // Erreur originale wrapped
}

// Error implémente l'interface error pour FileNotFoundError.
func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.FilePath)
}

// Unwrap permet le wrapping d'erreurs selon les conventions Go 1.13+.
func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

// ParseError représente une erreur de parsing lors de l'analyse d'un log.
// Elle contient l'identifiant du log et un message descriptif de l'erreur.
type ParseError struct {
	LogID   string // Identifiant du log en erreur
	Message string // Message descriptif de l'erreur de parsing
}

// Error implémente l'interface error pour ParseError.
func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing pour le log '%s': %s", e.LogID, e.Message)
}

// NewFileNotFoundError crée une nouvelle instance de FileNotFoundError.
func NewFileNotFoundError(filePath string, originalErr error) *FileNotFoundError {
	return &FileNotFoundError{
		FilePath: filePath,
		Err:      originalErr,
	}
}

// NewParseError crée une nouvelle instance de ParseError.
func NewParseError(logID, message string) *ParseError {
	return &ParseError{
		LogID:   logID,
		Message: message,
	}
}

// IsFileNotFoundError vérifie si une erreur est de type FileNotFoundError
// en utilisant errors.As() pour une détection robuste.
func IsFileNotFoundError(err error) bool {
	var fileErr *FileNotFoundError
	return errors.As(err, &fileErr)
}

// IsParseError vérifie si une erreur est de type ParseError
// en utilisant errors.As() pour une détection robuste.
func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}
