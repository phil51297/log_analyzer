# 📊 LogAnalyzer - Analyseur de Logs Parallèle

[![Go Version](https://img.shields.io/badge/go-1.24.3-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Un outil en ligne de commande (CLI) performant développé en Go pour analyser des fichiers de logs en parallèle. LogAnalyzer utilise la concurrence native de Go pour traiter simultanément plusieurs fichiers de logs et générer des rapports détaillés avec une gestion d'erreurs robuste.


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
