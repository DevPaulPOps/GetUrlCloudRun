# 🚀 GetUrlCloudRun

[![Go](https://img.shields.io/badge/Go-1.22.2-00ADD8.svg)](https://golang.org)
[![Gin](https://img.shields.io/badge/Gin-1.9.1-00ADD8.svg)](https://gin-gonic.com/)
[![Cloud Run](https://img.shields.io/badge/Cloud%20Run-1.3.6-4285F4.svg)](https://cloud.google.com/run)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED.svg)](https://www.docker.com/)
[![Terraform](https://img.shields.io/badge/Terraform-Ready-7B42BC.svg)](https://www.terraform.io)

Microservice Go permettant de récupérer les URLs des services Cloud Run déployés sur Google Cloud Platform via une API REST.

## 📑 Table des Matières

- [Objectifs](#-objectifs)
- [Architecture](#-architecture)
- [Stack Technique](#-stack-technique)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [API Reference](#-api-reference)
- [Infrastructure](#-infrastructure)
- [Contribution](#-contribution)
- [Authentification](#-authentification)

## 🎯 Objectifs

1. Authentification GCP sécurisée
2. Listing des services Cloud Run
3. Récupération des métadonnées de service
4. Exposition des URLs via API REST
5. Déploiement automatisé sur GCP

## 🏗️ Architecture

```sh
.
├── api/           # Handlers API REST et routes Gin
├── config/        # Configuration et gestion des paramètres
├── gcp/           # Client et services Google Cloud
├── model/         # Modèles de données
├── terraform/     # Infrastructure as Code
└── util/          # Utilitaires et gestion d'erreurs
```

## 💻 Stack Technique

### Core

- Go 1.22.2
- Gin Web Framework 1.9.1
- Cloud Run API 1.3.6

### Google Cloud Platform

- Cloud SDK v0.112.1
- Cloud Run
- Container Registry
- Secret Manager

### Dépendances Principales

- gin-gonic/gin v1.9.1
- gin-contrib/cors v1.7.1
- google.golang.org/api v0.169.0
- protobuf v1.33.0

### Infrastructure

- Terraform
- Docker

## 🚀 Installation

### Prérequis

- Go >= 1.22.2
- Google Cloud SDK
- Terraform >= 1.0
- Docker >= 20.10
- Compte GCP avec Cloud Run activé

### Installation Locale

```bash
# Clone du repository
git clone git@github.com:DevPaulPOps/GetUrlCloudRun.git
cd GetUrlCloudRun

# Installation des dépendances
go mod download
go mod tidy

# Vérification de la version Go
go version  # Doit afficher go1.22.2 ou supérieur

# Lancement du service
go run main.go
```

### Installation Docker

```bash
# Build de l'image
docker build -t geturlcloudrun:latest .

# Lancement avec configuration
docker run -p 3000:3000 \
  -v $(pwd)/config.json:/go/src/app/config.json \
  geturlcloudrun:latest
```

## ⚙️ Configuration

### Structure du fichier config.json

```json
{
	"AccountJSON": "/path/to/service-account.json",
	"Parent": "projects/YOUR_PROJECT/locations/REGION"
}
```

### Variables d'Environnement

```env
GCP_PROJECT_ID=your-project-id
GCP_REGION=us-central1
```

## 📡 API Reference

### Lister tous les services

```http
GET /service
```

#### Réponse

```json
[
	{
		"Name": "service-name",
		"Url": "https://service-xxx-uc.a.run.app"
	}
]
```

### Récupérer un service spécifique

```http
GET /service/:name
```

#### Réponse

```json
{
	"url": "https://service-xxx-uc.a.run.app"
}
```

## 🏗️ Infrastructure

### Terraform

```bash
cd terraform
terraform init
terraform plan
terraform apply
```

### Pipeline CI/CD (Cloud Build)

1. Build de l'image Docker
2. Push vers Artifact Registry
3. Déploiement Terraform
4. Configuration Cloud Run

Configuration détaillée dans `cloudbuild.yaml`

## 🔐 Authentification GCP

1. Créer un compte de service GCP

   - Accéder à la console GCP
   - IAM & Admin > Comptes de service
   - Créer un nouveau compte

2. Attribuer les rôles nécessaires

   - `roles/run.viewer` (minimum requis)
   - `roles/storage.objectViewer` (pour Artifact Registry)

3. Générer et télécharger les credentials
   - Format JSON
   - Configurer le chemin dans `config.json`

## 📝 Bonnes Pratiques

- Utilisation de contexts pour la gestion des timeouts
- Gestion d'erreurs centralisée
- Configuration externalisée
- Architecture modulaire
- Documentation complète des APIs

## 🤝 Contribution

1. Fork le projet
2. Créer une branche (`git checkout -b feature/amazing-feature`)
3. Commit (`git commit -m 'feat: Add amazing feature'`)
4. Push (`git push origin feature/amazing-feature`)
5. Ouvrir une Pull Request

---

_Développé avec ❤️ pour simplifier la gestion des services Cloud Run_
