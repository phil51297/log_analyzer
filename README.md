# 📊 LogAnalyzer - Analyseur de Logs Parallèle

[![Go Version](https://img.shields.io/badge/go-1.24.3-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Un outil en ligne de commande (CLI) performant développé en Go pour analyser des fichiers de logs en parallèle. LogAnalyzer utilise la concurrence native de Go pour traiter simultanément plusieurs fichiers de logs et générer des rapports détaillés avec une gestion d'erreurs robuste.

## 🎯 Fonctionnalités

- ✅ **Analyse parallèle** de multiples fichiers de logs via goroutines
- ✅ **Gestion d'erreurs robuste** avec types d'erreurs personnalisées
- ✅ **Interface CLI intuitive** basée sur Cobra
- ✅ **Import/Export JSON** pour configuration et rapports
- ✅ **Architecture modulaire** avec packages internes
- ✅ **Mode verbeux** pour le débogage
- ✅ **Simulation réaliste** d'analyse avec délais et erreurs aléatoires

## 🚀 Installation

### Prérequis

- Go 1.24.3 ou supérieur

### Compilation

```bash
# Cloner le repository
git clone https://github.com/phil51297/log_analyzer.git
cd log_analyzer

# Installer les dépendances
go mod tidy

# Compiler l'application
go build -o log_analyzer .

# Ou utiliser l'exécutable pré-compilé
# ./log_analyzer.exe (Windows)
# ./log_analyzer (Linux/macOS)
```

## 📖 Utilisation

### Commande de base

```bash
./log_analyzer --help
```

### Analyse de logs

```bash
# Analyse basique
./log_analyzer analyze --config config.json

# Analyse avec export et mode verbeux
./log_analyzer analyze -c config.json -o report.json -v
```

### Structure du fichier de configuration

Créez un fichier `config.json` avec la liste des logs à analyser :

```json
[
  {
    "id": "web-server-1",
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
    "type": "custom-app"
  }
]
```

### Exemple de rapport généré

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "/var/log/nginx/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "fichier introuvable: /non/existent/log.log"
  }
]
```

## 🏗️ Architecture

Le projet suit une architecture modulaire avec les packages suivants :

```
log_analyzer/
├── cmd/                    # Commandes CLI
│   ├── root.go            # Commande racine
│   └── analyze.go         # Commande d'analyse
├── internal/
│   ├── config/            # Gestion des configurations
│   │   └── config.go      # Lecture JSON
│   ├── analyzer/          # Moteur d'analyse
│   │   ├── analyzer.go    # Logique d'analyse
│   │   └── errors.go      # Erreurs personnalisées
│   └── reporter/          # Export des résultats
│       └── reporter.go    # Export JSON
├── test_logs/             # Fichiers de test
├── config.json            # Configuration exemple
└── main.go               # Point d'entrée
```

### Packages internes

#### `internal/config`

- **Responsabilité** : Chargement et validation des configurations JSON
- **Fonctions principales** :
  - `LoadConfig(configPath string)` : Charge la configuration depuis un fichier JSON

#### `internal/analyzer`

- **Responsabilité** : Analyse des logs et gestion des erreurs
- **Fonctions principales** :
  - `AnalyzeLog(logConfig config.LogConfig)` : Analyse un fichier de log
  - Erreurs personnalisées : `FileNotFoundError`, `ParseError`
  - Utilitaires : `IsFileNotFoundError()`, `IsParseError()`

#### `internal/reporter`

- **Responsabilité** : Export des résultats
- **Fonctions principales** :
  - `ExportToJSON(results []analyzer.AnalysisResult, filePath string)` : Export vers fichier JSON

## 🔧 Concepts techniques implémentés

### Concurrence

- **Goroutines** : Une goroutine par fichier de log à analyser
- **WaitGroup** : Synchronisation des goroutines
- **Channels** : Collecte sécurisée des résultats

### Gestion d'erreurs

- **Erreurs personnalisées** : `FileNotFoundError` et `ParseError`
- **Error wrapping** : Utilisation d'`errors.As()` et `errors.Is()`
- **Messages localisés** : Erreurs en français avec détails techniques

### CLI avec Cobra

- **Commandes structurées** : Commande racine et sous-commandes
- **Flags typés** : `-c/--config`, `-o/--output`, `-v/--verbose`
- **Validation** : Flag config marqué comme requis

## 🎲 Simulation d'analyse

L'outil simule une analyse réaliste avec :

- **Délai aléatoire** : 50-200ms par fichier (simulation de traitement)
- **Erreurs aléatoires** : 10% de chance d'erreur de parsing
- **Vérifications** : Existence, lisibilité et type de fichier

## 📝 Exemples d'utilisation

### Analyse simple

```bash
./log_analyzer analyze -c config.json
```

### Analyse avec export

```bash
./log_analyzer analyze -c config.json -o my_report.json
```

### Mode verbeux pour débogage

```bash
./log_analyzer analyze -c config.json -v
```

### Fichiers de test inclus

```bash
# Utiliser les fichiers de test inclus
./log_analyzer analyze -c config.json -o test_report.json -v
```

## 👥 Équipe de développement

- **Développeur principal** : phil51297
- **Framework utilisé** : Go + Cobra CLI
- **Version** : 1.0.0

## 🔍 Tests et validation

Le projet inclut :

- ✅ Validation manuelle avec fichiers de test
- ✅ Gestion de cas d'erreur (fichiers inexistants)
- ✅ Test de concurrence avec multiples fichiers
- ✅ Validation des formats JSON

## 📋 Dépendances

- **Go** : 1.24.3+
- **github.com/spf13/cobra** : v1.9.1 (CLI framework)

## 🚦 Status du projet

Le projet respecte entièrement le cahier des charges avec :

- ✅ Toutes les fonctionnalités demandées implémentées
- ✅ Architecture modulaire respectée
- ✅ Concurrence et gestion d'erreurs conformes
- ✅ Interface CLI complète et fonctionnelle
- ✅ Documentation complète

## 📄 Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.
