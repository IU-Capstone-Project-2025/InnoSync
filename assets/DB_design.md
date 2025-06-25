

# InnoSync Project – Database Design Documentation

## Table of Contents
- [InnoSync Project – Database Design Documentation](#innosync-project--database-design-documentation)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Core Tables](#core-tables)
    - [USER](#user)
    - [REFRESH\_TOKEN](#refresh_token)
    - [USER\_PROFILE](#user_profile)
    - [WORK\_EXPERIENCE](#work_experience)
    - [TECHNOLOGY](#technology)
    - [USER\_PROFILE\_TECHNOLOGY](#user_profile_technology)
  - [Project \& Recruitment](#project--recruitment)
    - [PROJECT](#project)
    - [PROJECT\_ROLE](#project_role)
    - [INVITATION](#invitation)
  - [FastSyncing (AI Matching)](#fastsyncing-ai-matching)
    - [FASTSYNC\_REQUEST](#fastsync_request)
    - [FASTSYNC\_SUGGESTION](#fastsync_suggestion)
  - [Collaboration \& Communication](#collaboration--communication)
    - [GROUP\_CHAT](#group_chat)
    - [GROUP\_CHAT\_MEMBER](#group_chat_member)
    - [MESSAGE](#message)
  - [Enum Types](#enum-types)
  - [Relationships Summary](#relationships-summary)
  - [Security Notes](#security-notes)

---

## Overview

The **InnoSync** database schema supports user authentication, detailed talent profiles, project recruitment workflows, AI-powered team matching (FastSyncing), and integrated collaboration tools.  
It is designed for flexibility, data integrity, and extensibility to scale with platform features.

---

## Core Tables

### USER

Stores basic authentication data for all users.

| Field       | Type      | Description                 |
|-------------|-----------|-----------------------------|
| id          | bigint    | Primary key                 |
| email       | string    | Unique email                |
| full_name   | string    | Full name                   |
| password   | string    | Hashed password             |
| created_at  | timestamp | Registration timestamp      |

---

### REFRESH_TOKEN

Stores JWT refresh tokens linked to users.

| Field       | Type      | Description                      |
|-------------|-----------|----------------------------------|
| id          | bigint    | Primary key                      |
| user_id     | bigint    | Foreign key → USER(id)           |
| token       | string    | Refresh token                    |
| expiry_date | timestamp | Expiration date                  |

---

### USER_PROFILE

Stores detailed user profile info beyond authentication.

| Field         | Type      | Description                                |
|---------------|-----------|--------------------------------------------|
| id            | bigint    | Primary key                                |
| user_id       | bigint    | Foreign key → USER(id)                      |
| telegram      | string    | Telegram handle (optional)                  |
| github        | string    | GitHub profile (optional)                   |
| bio           | string    | Short biography                             |
| position      | string    | Current position                            |
| education     | education_enum | Enum: NO_DEGREE, BACHELOR, MASTER, PHD  |
| expertise     | string    | Expertise areas (free text or tags)         |
| expertise_level | expertise_level_enum | Enum: ENTRY, JUNIOR, MID, SENIOR, RESEARCHER |
| resume        | string    | URL or path to uploaded resume              |

---

### WORK_EXPERIENCE

Records professional experiences for user profiles.

| Field           | Type      | Description                        |
|-----------------|-----------|----------------------------------|
| id              | bigint    | Primary key                      |
| user_profile_id  | bigint    | Foreign key → USER_PROFILE(id)   |
| start_date      | date      | Employment start date             |
| end_date        | date      | Employment end date (nullable)    |
| position        | string    | Job title                        |
| company         | string    | Name of company                   |
| description     | string    | Job description                   |

---

### TECHNOLOGY

Catalog of technologies used for tagging user skills and project needs.

| Field | Type   | Description      |
|-------|--------|------------------|
| id    | bigint | Primary key      |
| name  | string | Technology name  |

---

### USER_PROFILE_TECHNOLOGY

Join table for many-to-many relationship between user profiles and technologies.

| Field           | Type   | Description                    |
|-----------------|--------|--------------------------------|
| user_profile_id | bigint | Foreign key → USER_PROFILE(id) |
| technology_id   | bigint | Foreign key → TECHNOLOGY(id)   |

---

## Project & Recruitment

### PROJECT

Stores projects created by recruiters.

| Field           | Type      | Description                      |
|-----------------|-----------|----------------------------------|
| id              | bigint    | Primary key                      |
| recruiter_id    | bigint    | Foreign key → USER(id)           |
| title           | string    | Project title                   |
| description     | string    | Project description             |
| team_size       | integer   | Number of required team members |
| created_at      | timestamp | Project creation time           |

---

### PROJECT_ROLE

Defines required roles and expectations per project.

| Field           | Type      | Description                     |
|-----------------|-----------|--------------------------------|
| id              | bigint    | Primary key                    |
| project_id      | bigint    | Foreign key → PROJECT(id)       |
| role_name       | string    | Role name (e.g., Frontend Dev) |
| expertise_level | expertise_level_enum | Required expertise level    |
| technologies    | string    | Technologies required (CSV or separate join table possible) |

---

### INVITATION

Tracks project invitations sent from recruiters to recruitees.

| Field           | Type      | Description                      |
|-----------------|-----------|---------------------------------|
| id              | bigint    | Primary key                     |
| project_role_id | bigint    | Foreign key → PROJECT_ROLE(id)  |
| user_id         | bigint    | Foreign key → USER(id)           |
| status          | invitation_status_enum | INVITED, ACCEPTED, DECLINED, REVOKED |
| sent_at         | timestamp | Invitation send time             |
| responded_at    | timestamp | Time user responded              |

---
### APPLICATION (NEW)
Tracks user applications to project roles.
| Field           | Type      | Description                      |
|-----------------|-----------|---------------------------------|
| id              | bigint    | Primary key                     |
| user_id | bigint    | Foreign key → User(id)  |
| project_role_id         | bigint    | Foreign key → Project_role(id)           |
| status | application_status_enum | PENDING, UNDER_REVIEW, ACCEPTED, REJECTED, WITHDRAWN |
| applied_at | timestamp | Application submission time          |
| updated_at | timestamp | Last status update time |

## FastSyncing (AI Matching)

### FASTSYNC_REQUEST

Represents a recruiter opting into AI-powered matching.

| Field           | Type      | Description                      |
|-----------------|-----------|---------------------------------|
| id              | bigint    | Primary key                     |
| project_id      | bigint    | Foreign key → PROJECT(id)        |
| requested_at    | timestamp | Request timestamp               |

---

### FASTSYNC_SUGGESTION

AI-generated suggestions of candidates for specific project roles.

| Field             | Type      | Description                      |
|-------------------|-----------|---------------------------------|
| id                | bigint    | Primary key                     |
| fastsync_request_id | bigint   | Foreign key → FASTSYNC_REQUEST(id) |
| project_role_id   | bigint    | Foreign key → PROJECT_ROLE(id)   |
| suggested_user_id | bigint    | Foreign key → USER(id)           |
| status            | suggestion_status_enum | PENDING, REJECTED, APPROVED   |

---

## Collaboration & Communication

### GROUP_CHAT

Stores group chat info for project teams.

| Field       | Type      | Description                 |
|-------------|-----------|-----------------------------|
| id          | bigint    | Primary key                 |
| project_id  | bigint    | Foreign key → PROJECT(id)    |
| created_at  | timestamp | Group chat creation time   |

---

### GROUP_CHAT_MEMBER

Maps users to group chats.

| Field         | Type      | Description                 |
|---------------|-----------|-----------------------------|
| group_chat_id | bigint    | Foreign key → GROUP_CHAT(id) |
| user_id       | bigint    | Foreign key → USER(id)       |

---

### MESSAGE

Stores messages for group chats and direct messages.

| Field        | Type      | Description                       |
|--------------|-----------|---------------------------------|
| id           | bigint    | Primary key                     |
| sender_id    | bigint    | Foreign key → USER(id)           |
| group_chat_id| bigint    | Foreign key → GROUP_CHAT(id) (nullable for DMs) |
| content      | string    | Message text                    |
| sent_at      | timestamp | Message sent time               |

---

## Enum Types

- **education_enum**: `NO_DEGREE`, `BACHELOR`, `MASTER`, `PHD`  
- **expertise_level_enum**: `ENTRY`, `JUNIOR`, `MID`, `SENIOR`, `RESEARCHER`  
- **invitation_status_enum**: `INVITED`, `ACCEPTED`, `DECLINED`, `REVOKED`  
- **suggestion_status_enum**: `PENDING`, `REJECTED`, `APPROVED`  

---

## Relationships Summary

| Relationship                         | Cardinality          |
|------------------------------------|----------------------|
| USER → REFRESH_TOKEN                | One-to-Many          |
| USER → USER_PROFILE                 | One-to-One           |
| USER_PROFILE → WORK_EXPERIENCE     | One-to-Many          |
| USER_PROFILE → TECHNOLOGY           | Many-to-Many         |
| USER (Recruiter) → PROJECT          | One-to-Many          |
| PROJECT → PROJECT_ROLE              | One-to-Many          |
| PROJECT_ROLE → INVITATION           | One-to-Many          |
| USER → APPLICATION                   | One-to-Many          |
| PROJECT_ROLE → APPLICATION           | One-to-Many          |
| PROJECT → GROUP_CHAT                | One-to-One           |
| GROUP_CHAT → GROUP_CHAT_MEMBER     | One-to-Many          |
| USER → MESSAGE                     | One-to-Many          |
| GROUP_CHAT → MESSAGE               | One-to-Many (nullable) |

---

## Security Notes

- Passwords are stored hashed using secure algorithms (e.g., bcrypt).  
- Tokens in REFRESH_TOKEN must be securely generated and stored.  
- Enum sets are locked down at the database level for data integrity.  
- Sensitive profile links (GitHub, LinkedIn) should be validated.  
- Role-based access control is implemented at the backend layer.

---
