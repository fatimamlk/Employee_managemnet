# Employee Management API

Une API RESTful de gestion des employés développée avec **Go** (Golang) et **MongoDB**. Cette API permet d'effectuer des opérations CRUD (Créer, Lire, Mettre à jour, Supprimer) sur les informations des employés.

## Fonctionnalités principales

- **Lister les employés** : Récupérer la liste de tous les employés.
- **Ajouter un employé** : Ajouter un nouvel employé à la base.
- **Modifier un employé** : Mettre à jour les informations d'un employé existant.
- **Supprimer un employé** : Supprimer un employé de la base de données.

## Endpoints de l'API

### Base URL : `http://localhost:8080`

| Méthode | Endpoint            | Description                    |
|---------|---------------------|--------------------------------|
| `GET`   | `/employees`         | Récupère tous les employés     |
| `GET`   | `/employees/{id}`    | Récupère un employé par ID     |
| `POST`  | `/employees`         | Ajoute un nouvel employé       |
| `PUT`   | `/employees/{id}`    | Met à jour un employé existant |
| `DELETE`| `/employees/{id}`    | Supprime un employé par ID     |

## Installation

### Prérequis

- **Go** (Golang) 1.16 ou supérieur
- **MongoDB** installé et en fonctionnement sur votre machine ou accessible via une instance distante.
- **Go MongoDB driver** : Vous devrez installer le driver MongoDB pour Go.

### Étapes d'installation

1. **Cloner le repository** :
   ```bash
   git clone https://github.com/yourusername/employee-management-api.git
   cd employee-management-api
### Installer les dépendances

Si vous utilisez **go mod**, assurez-vous d'initialiser votre module Go avec les commandes suivantes :

```bash
go mod init employee-management-api
go mod tidy
