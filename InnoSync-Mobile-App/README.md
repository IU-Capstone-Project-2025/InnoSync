# InnoSync Mobile App

A full-stack, collaborative project management platform for innovators, teams, and talent. This monorepo contains both the backend RESTful API (Go, Gin, PostgreSQL) and the cross-platform mobile frontend (Flutter).

---

## Project Overview

**InnoSync:** InnoSync & team-matching platform connecting talent and projects in university communities.
**Features:**
- Register, authenticate, and manage profiles
- Create and manage projects and roles
- Apply for project roles and track applications
- Send and respond to invitations

---

## Architecture

```
InnoSyncMobileApp/
├── backend/   # Go, Gin, PostgreSQL REST API
└── frontend/  # Flutter mobile app (Android, iOS, Web, Desktop)
```

- **Backend:** Go, Gin, PostgreSQL, Docker
- **Frontend:** Flutter (Dart)
- **DevOps:** Github Actions CI/CD
---

## Features

### Backend
- JWT authentication (access/refresh tokens)
- User profile management (with file uploads)
- Project and role management
- Application and invitation system
- RESTful API, CORS, Dockerized

### Frontend
- Modern Flutter UI
- User login/signup
- Project discovery and application
- Profile creation and editing
- Team and invitation management

---

## Project Structure

```
backend/   # API, business logic, database, Docker
frontend/  # Flutter app, widgets, screens, models
```

---

## Setup Instructions

### Prerequisites
- **Backend:** Go 1.21+, PostgreSQL 12+, Docker & Docker Compose (optional)
- **Frontend:** Flutter SDK 3.8+, Dart 3.8+

### Backend Setup

#### 1. Local Development

- Install Go and PostgreSQL
- Copy `backend/config.yaml` and edit as needed:
  ```yaml
  server:
    port: 3000
    host: "0.0.0.0"
    mode: "debug"
  database:
    host: "localhost"
    port: 5432
    user: "postgres"
    password: "postgres"
    dbname: "innosync"
    sslmode: "disable"
  jwt:
    secret_key: "your-secret-key-change-this-in-production"
    access_token_ttl: "15m"
    refresh_token_ttl: "168h"
  upload:
    max_file_size: 10485760
    upload_dir: "uploads"
    allowed_extensions: [".jpg", ".jpeg", ".png", ".pdf", ".doc", ".docx"]
  ```
- Install dependencies:
  ```bash
  cd backend
  go mod download
  ```
- Run database migrations (if any)
- Start the server:
  ```bash
  go run main.go
  ```
- API available at `http://localhost:3000/api/v1`

#### 2. Docker Compose (Recommended)

- Ensure Docker and Docker Compose are installed
- From the project root:
  ```bash
  docker-compose up -d
  ```
- Services:
  - **db:** PostgreSQL (port 5432)
  - **backend:** Go API (port 3000)
  - **adminer:** DB UI (port 3001)

---

### Frontend Setup

- Install [Flutter SDK](https://docs.flutter.dev/get-started/install)
- From the `frontend` directory:
  ```bash
  flutter pub get
  flutter run
  ```
- To run on web/desktop, see [Flutter docs](https://docs.flutter.dev/)

#### Linting & Code Quality
- Uses `flutter_lints` (see `analysis_options.yaml`)
- Run:
  ```bash
  flutter analyze
  ```

## 📄 API Documentation

Access and interact with our API seamlessly:

- **Swagger UI** – Explore and test endpoints interactively.
- **Markdown Guide** – Full API spec in [backend/README.md](https://github.com/team-18-project/InnoSync/blob/main/backend/README.md).

### 🔍 API Highlights

- **Authentication**
  - JWT based with secure login and token refresh.
- **Core Endpoints**
  - **Users**: register, login, profile management.
  - **Projects**: create, read, update, delete; team invites.
---

## 🗄️ Database Schema & Design

Visual and descriptive overview available:

- **Schema Diagram**
  ![Database Schema](https://raw.githubusercontent.com/team-18-project/InnoSync/refs/heads/main/db/DB_diagram.png)
- **Detailed ERD & Tables** – See [db/DB_design.md](https://github.com/team-18-project/InnoSync/blob/main/db/DB_design.md)
  - Tables: `users`, `projects`, `tasks`, `comments`, `notifications`, `activities`, etc.
  - Relationships: one-to-many (users→projects, projects→tasks), many-to-many (users↔projects via project_team join table).

---

---

## 🧠 Product Design & UX Research

Comprehensive user-research-backed features:

- In-depth findings and validation in [Product Design & UX Research](https://www.figma.com/design/islrijonPzZEUrBi9NgV94/UX-UI-Mid?node-id=47-31994&t=bATNSWtlbjcDZmov-1)
- Highlights:
  - **User Pain Points** – Collaboration friction, task visibility gaps.
  - **Proposed Solutions** – Real-time updates, shared project dashboards, invitation workflows.
  - **Feature Roadmap** – Prioritized through user feedback and impact analysis.
  - **Wireframes & Mockups** – Visual representations of key screens and user flows.
- High-Fideliy Design [Link](https://www.figma.com/design/qDq1rHa8D5f1AlJtMBb0rz/High-Fi-design-InnoSync?t=5BQsZTzFq7z2nK77-0)

---

## Contributions

- Stefan Farafonov: Team lead and management, Flutter Frontend Implementation
- Ahmed Baha Eddine Alimi: Figma Design, API implementation and documentation, Product Conception
- Aigul Farkhetdinovaa: Database Implementation, CI/CD pipeline setup, Unit Tests Implementation
- Arifzhan Narimov: Flutter Frontend Implementation
- Roman Voronin: Endpoint Connection

---

## 🌐 Web Application Access

In addition to the mobile app, we also offer a **Web Application** with full project management functionalities.
You can access it here: [https://innosync.duckdns.org](https://innosync.duckdns.org)

The web app ensures a seamless experience across both platforms, with plans for advanced account syncing and unified notifications.

---

## 🚀 Future Plans

- **Quick Syncing (AI-Powered Smart Matching)**
  - AI-driven recommendations for team matching and project collaborations.

- **Cross-Platform Account Syncing**
  - Seamless integration between mobile and web applications for unified access and synchronized notifications.

--- 
# Checklist  
## Technical requirements (20 points)

### Backend Development (8 points)
- [x] Gobased backend (3 points)
- [x] RESTful API with Swagger documentation (2 points)
- [x] PostgreSQL database with proper schema design (1 point)
- [x] JWT-based authentication and authorization (1 point)
- [x] Comprehensive unit and integration tests (1 point)

### Frontend Development (8 points)
- [x] Flutter-based cross-platform application (mobile + web) (3 points)
- [x] Responsive UI design with custom widgets (1 point)
- [x] State management implementation (1 point)
- [ ] Offline data persistence (1 point)
- [x] Unit and widget tests (1 point)
- [x] Support light and dark mode (1 point)

### DevOps & Deployment (4 points)
- [x] Docker compose for all services (1 point)
- [x] CI/CD pipeline implementation (1 point)
- [x] Environment configuration management using config files (1 point)
- [x] GitHub pages for the project (1 point)

---
## Non-Technical Requirements (10 points)

### Project Management (4 points)
- [x] GitHub organization with well-maintained repository (1 point)
- [x] Regular commits and meaningful pull requests from all team members (1 point)
- [x] Project board (GitHub Projects) with task tracking (1 point)
- [x] Team member roles and responsibilities documentation (1 point)

### Documentation (4 points)
-  Comprehensive README with:
  - [x] Project overview and setup instructions (1 point)
  - [ ] Screenshots and GIFs of key features (1 point)
  - [x] API documentation (1 point)
  - [x] Architecture diagrams and explanations (1 point)

### Code Quality (2 points)
- [x] Consistent code style and formatting during CI/CD pipeline (1 point)
- [x] Code review participation and resolution (1 point)

## Bonus points (up to 10 points total)
- [ ] Implement localization for Russian (RU) and English (ENG) languages (2 points)
- [x] Good UI/UX design (up to 3 points) - subjective
- [ ] Integration with external APIs (fitness trackers, health devices) (up to 5 points)
- [x] Comprehensive error handling and user feedback (up to 2 points)
- [ ] Advanced animations and transitions (up to 3 points)
- [ ] Widget implementation for native mobile elements (up to 2 points)  
Total: 28 up to 33 with bonuses (if we get bonuses :))

---
## License

MIT License
