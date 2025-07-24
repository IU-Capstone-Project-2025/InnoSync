# InnoSync Mobile App Backend

A RESTful API backend for the InnoSync mobile application built with Go, Gin framework, and PostgreSQL.

## Features

### 🔐 Authentication
- User registration and login
- JWT access and refresh tokens
- Token refresh mechanism
- Secure password hashing with bcrypt
- User logout functionality

### 👤 Profile Management
- Create and manage user profiles
- File upload support (profile images, resumes)
- Work experience management
- Technology/skills management
- Education and expertise level tracking

### 🚀 Project Management
- Create and manage projects
- Project role management
- Team size tracking
- Project visibility controls

### 📝 Application System
- Apply for project roles
- Application status tracking
- Cover letter support
- Application withdrawal

### 💌 Invitation System
- Send invitations to project roles
- Respond to invitations
- Revoke invitations
- Track invitation status

### 🛠️ Infrastructure
- PostgreSQL database
- RESTful API with Gin framework
- Configuration management
- File upload support
- CORS middleware
- Docker containerization
- Health check endpoint

## Project Structure

```
backend/
├── cmd/                 # Application entry points
├── config/             # Configuration management
├── handlers/           # HTTP handlers
├── middleware/         # HTTP middleware
├── models/            # Data models
├── services/          # Business logic
├── repo/              # Data access layer
├── utils/             # Utility functions
├── uploads/           # File uploads directory
├── Dockerfile         # Docker configuration
├── config.yaml        # Application configuration
├── go.mod             # Go modules
└── main.go            # Main application file
```

## Setup and Installation

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Docker & Docker Compose (optional)

### Running with Docker

1. Clone the repository
2. Copy `config.yaml.example` to `config.yaml` and configure your settings
3. Run with Docker Compose:
```bash
docker-compose up -d
```

### Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Configure database connection in `config.yaml`

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
Most endpoints require authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

---

## 🔐 Authentication APIs

### Register User
**POST** `/auth/signup`

Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123",
  "full_name": "John Doe"
}
```

**Response (201 Created):**
```json
{
  "user": {
    "id": 1,
    "email": "user@example.com",
    "full_name": "John Doe",
    "created_at": "2024-01-01T00:00:00Z"
  },
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 3600
  }
}
```

### Login
**POST** `/auth/login`

Authenticate a user and return tokens.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**
```json
{
  "user": {
    "id": 1,
    "email": "user@example.com",
    "full_name": "John Doe",
    "created_at": "2024-01-01T00:00:00Z",
    "last_login": "2024-01-01T12:00:00Z"
  },
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 3600
  }
}
```

### Refresh Token
**POST** `/auth/refresh`

Refresh the access token using a valid refresh token.

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

### Logout
**POST** `/auth/logout`

Invalidate the refresh token.

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (200 OK):**
```json
{
  "message": "Successfully logged out"
}
```

### Get Current User
**GET** `/auth/me`

Get information about the currently authenticated user.

**Response (200 OK):**
```json
{
  "id": 1,
  "email": "user@example.com",
  "full_name": "John Doe",
  "created_at": "2024-01-01T00:00:00Z",
  "last_login": "2024-01-01T12:00:00Z"
}
```

---

## 👤 Profile APIs

### Create Profile
**POST** `/profile`

Create a new user profile with detailed information.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "telegram": "@johndoe",
  "github": "https://github.com/johndoe",
  "bio": "Experienced software developer",
  "position": "Senior Software Engineer",
  "education": "MASTER",
  "expertise": "Backend Development",
  "expertise_level": "SENIOR",
  "technologies": ["Go", "PostgreSQL", "Docker"],
  "work_experience": [
    {
      "start_date": "2020-01-01",
      "end_date": "2023-12-31",
      "position": "Software Engineer",
      "company": "Tech Corp",
      "description": "Developed microservices"
    }
  ]
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "telegram": "@johndoe",
  "github": "https://github.com/johndoe",
  "bio": "Experienced software developer",
  "position": "Senior Software Engineer",
  "education": "MASTER",
  "expertise": "Backend Development",
  "expertise_level": "SENIOR",
  "work_experiences": [
    {
      "id": 1,
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2023-12-31T00:00:00Z",
      "position": "Software Engineer",
      "company": "Tech Corp",
      "description": "Developed microservices"
    }
  ],
  "technologies": [
    {"id": 1, "name": "Go"},
    {"id": 2, "name": "PostgreSQL"},
    {"id": 3, "name": "Docker"}
  ],
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Get User Profile
**GET** `/profile/me`

Get the authenticated user's profile information.

**Response (200 OK):**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "telegram": "@johndoe",
  "github": "https://github.com/johndoe",
  "bio": "Experienced software developer",
  "position": "Senior Software Engineer",
  "education": "MASTER",
  "expertise": "Backend Development",
  "expertise_level": "SENIOR",
  "resume_url": "/uploads/resumes/resume.pdf",
  "avatar_url": "/uploads/images/avatar.jpg",
  "work_experiences": [...],
  "technologies": [...],
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Update Profile
**PUT** `/profile`

Update user profile information.

**Request Body:**
```json
{
  "telegram": "@newtelegram",
  "bio": "Updated bio",
  "position": "Lead Developer",
  "expertise_level": "SENIOR"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "user_id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "telegram": "@newtelegram",
  "bio": "Updated bio",
  "position": "Lead Developer",
  "expertise_level": "SENIOR",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Upload Profile Image
**POST** `/profile/image`

Upload a profile image for the authenticated user.

**Request Body:** `multipart/form-data`
- `image`: Image file (JPG, PNG, GIF)

**Response (200 OK):**
```json
{
  "url": "/uploads/images/12345-abc123.jpg",
  "filename": "12345-abc123.jpg",
  "size": 1024768,
  "message": "Profile image uploaded successfully"
}
```

### Upload Resume
**POST** `/profile/resume`

Upload a resume file for the authenticated user.

**Request Body:** `multipart/form-data`
- `resume`: Resume file (PDF, DOC, DOCX)

**Response (200 OK):**
```json
{
  "url": "/uploads/resumes/12345-resume.pdf",
  "filename": "12345-resume.pdf",
  "size": 2048576,
  "message": "Resume uploaded successfully"
}
```

### Get Technologies
**GET** `/profile/technologies`

Get list of all available technologies.

**Response (200 OK):**
```json
[
  {"id": 1, "name": "Go", "category": "Backend"},
  {"id": 2, "name": "JavaScript", "category": "Frontend"},
  {"id": 3, "name": "PostgreSQL", "category": "Database"}
]
```

---

## 🚀 Project APIs

### Create Project
**POST** `/projects`

Create a new project.

**Request Body:**
```json
{
  "title": "InnoSync Mobile App",
  "description": "A collaborative project management app",
  "team_size": 5
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "title": "InnoSync Mobile App",
  "description": "A collaborative project management app",
  "team_size": 5,
  "is_active": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Get User Projects
**GET** `/projects/me`

Get all projects created by the authenticated user.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "title": "InnoSync Mobile App",
    "description": "A collaborative project management app",
    "team_size": 5,
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "roles": [
      {
        "id": 1,
        "role_name": "Backend Developer",
        "expertise_level": "MID",
        "technologies": ["Go", "PostgreSQL"],
        "is_open": true
      }
    ]
  }
]
```

### Get Project
**GET** `/projects/{id}`

Get a specific project by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "title": "InnoSync Mobile App",
  "description": "A collaborative project management app",
  "team_size": 5,
  "is_active": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Update Project
**PUT** `/projects/{id}`

Update a project.

**Request Body:**
```json
{
  "title": "Updated Project Title",
  "description": "Updated description",
  "team_size": 10
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "title": "Updated Project Title",
  "description": "Updated description",
  "team_size": 10,
  "is_active": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Delete Project
**DELETE** `/projects/{id}`

Delete a project (soft delete).

**Response (204 No Content)**

### Create Project Role
**POST** `/projects/{id}/roles`

Create a new role for a project.

**Request Body:**
```json
{
  "role_name": "Frontend Developer",
  "expertise_level": "MID",
  "technologies": ["React", "TypeScript", "CSS"]
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "project_id": 1,
  "role_name": "Frontend Developer",
  "expertise_level": "MID",
  "technologies": ["React", "TypeScript", "CSS"],
  "is_open": true,
  "project_title": "InnoSync Mobile App"
}
```

### Get Project Roles
**GET** `/projects/{id}/roles`

Get all roles for a specific project.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "project_id": 1,
    "role_name": "Frontend Developer",
    "expertise_level": "MID",
    "technologies": ["React", "TypeScript", "CSS"],
    "is_open": true,
    "project_title": "InnoSync Mobile App"
  }
]
```

### Get Project Role
**GET** `/projects/{id}/roles/{roleId}`

Get a specific project role by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "project_id": 1,
  "role_name": "Frontend Developer",
  "expertise_level": "MID",
  "technologies": ["React", "TypeScript", "CSS"],
  "is_open": true,
  "project_title": "InnoSync Mobile App"
}
```

---

## 📝 Application APIs

### Create Application
**POST** `/applications/project-roles/{projectRoleId}`

Apply for a project role.

**Request Body:**
```json
{
  "cover_letter": "I am interested in this role because..."
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "user_id": 1,
  "user_full_name": "John Doe",
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "status": "PENDING",
  "applied_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "cover_letter": "I am interested in this role because..."
}
```

### Get User Applications
**GET** `/applications`

Get all applications submitted by the authenticated user.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_full_name": "John Doe",
    "project_role_id": 1,
    "role_name": "Frontend Developer",
    "project_id": 1,
    "project_title": "InnoSync Mobile App",
    "status": "PENDING",
    "applied_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Role Applications
**GET** `/applications/project-roles/{projectRoleId}`

Get all applications for a specific project role.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_full_name": "John Doe",
    "project_role_id": 1,
    "role_name": "Frontend Developer",
    "project_id": 1,
    "project_title": "InnoSync Mobile App",
    "status": "PENDING",
    "applied_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Application
**GET** `/applications/{id}`

Get a specific application by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "user_id": 1,
  "user_full_name": "John Doe",
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "status": "PENDING",
  "applied_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "cover_letter": "I am interested in this role because..."
}
```

### Update Application Status
**PATCH** `/applications/{id}/status`

Update the status of a specific application.

**Request Body:**
```json
{
  "status": "ACCEPTED"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "user_id": 1,
  "user_full_name": "John Doe",
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "status": "ACCEPTED",
  "applied_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Delete Application
**DELETE** `/applications/{id}`

Withdraw a submitted application.

**Response (204 No Content)**

---

## 💌 Invitation APIs

### Create Invitation
**POST** `/invitations`

Create a new invitation for a project role.

**Request Body:**
```json
{
  "project_role_id": 1,
  "recipient_id": 2,
  "message": "We would like to invite you to join our project"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "recipient_id": 2,
  "recipient_name": "Jane Smith",
  "status": "INVITED",
  "sent_at": "2024-01-01T00:00:00Z",
  "message": "We would like to invite you to join our project"
}
```

### Get Sent Invitations
**GET** `/invitations/sent`

Get all invitations sent by the authenticated user.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "project_role_id": 1,
    "role_name": "Frontend Developer",
    "project_id": 1,
    "project_title": "InnoSync Mobile App",
    "recipient_id": 2,
    "recipient_name": "Jane Smith",
    "status": "INVITED",
    "sent_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Received Invitations
**GET** `/invitations/received`

Get all invitations received by the authenticated user.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "project_role_id": 1,
    "role_name": "Frontend Developer",
    "project_id": 1,
    "project_title": "InnoSync Mobile App",
    "recipient_id": 2,
    "recipient_name": "Jane Smith",
    "status": "INVITED",
    "sent_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Invitation
**GET** `/invitations/{id}`

Get a specific invitation by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "recipient_id": 2,
  "recipient_name": "Jane Smith",
  "status": "INVITED",
  "sent_at": "2024-01-01T00:00:00Z",
  "message": "We would like to invite you to join our project"
}
```

### Respond to Invitation
**PATCH** `/invitations/{id}/respond`

Respond to an invitation.

**Request Body:**
```json
{
  "response": "ACCEPTED"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "recipient_id": 2,
  "recipient_name": "Jane Smith",
  "status": "ACCEPTED",
  "sent_at": "2024-01-01T00:00:00Z",
  "responded_at": "2024-01-01T12:00:00Z"
}
```

### Revoke Invitation
**PATCH** `/invitations/{id}/revoke`

Revoke a sent invitation.

**Response (200 OK):**
```json
{
  "id": 1,
  "project_role_id": 1,
  "role_name": "Frontend Developer",
  "project_id": 1,
  "project_title": "InnoSync Mobile App",
  "recipient_id": 2,
  "recipient_name": "Jane Smith",
  "status": "REVOKED",
  "sent_at": "2024-01-01T00:00:00Z",
  "responded_at": "2024-01-01T12:00:00Z"
}
```

### Delete Invitation
**DELETE** `/invitations/{id}`

Delete an invitation.

**Response (204 No Content)**

### Get Invitations by Project Role
**GET** `/invitations/project-roles/{projectRoleId}`

Get all invitations for a specific project role.

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "project_role_id": 1,
    "role_name": "Frontend Developer",
    "project_id": 1,
    "project_title": "InnoSync Mobile App",
    "recipient_id": 2,
    "recipient_name": "Jane Smith",
    "status": "INVITED",
    "sent_at": "2024-01-01T00:00:00Z"
  }
]
```

---

## 🏥 Health Check API

### Health Check
**GET** `/health`

Check the health status of the API.

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

---

## Data Models

### Enums

#### Education Level
- `NO_DEGREE`
- `BACHELOR`
- `MASTER`
- `PHD`

#### Expertise Level
- `ENTRY`
- `JUNIOR`
- `MID`
- `SENIOR`
- `RESEARCHER`

#### Application Status
- `PENDING`
- `UNDER_REVIEW`
- `ACCEPTED`
- `REJECTED`
- `WITHDRAWN`

#### Invitation Status
- `INVITED`
- `ACCEPTED`
- `DECLINED`
- `REVOKED`

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

Common HTTP status codes:
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required or invalid
- `403 Forbidden`: Access denied
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource already exists
- `500 Internal Server Error`: Server error

## Environment Variables

Create a `config.yaml` file with the following structure:

```yaml
server:
  host: localhost
  port: 8080
  mode: debug

database:
  host: localhost
  port: 5432
  name: innosync_db
  user: postgres
  password: password
  sslmode: disable

jwt:
  secret: your-secret-key
  access_token_duration: 1h
  refresh_token_duration: 168h

upload:
  upload_dir: ./uploads
  max_file_size: 10485760
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.