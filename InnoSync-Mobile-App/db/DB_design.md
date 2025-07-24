

# InnoSync Project – Database Design Documentation

## Table of Contents
- [InnoSync Project – Database Design Documentation](#innosync-project--database-design-documentation)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Core Tables](#core-tables)
    - [USERS](#users)
    - [REFRESH_TOKEN](#refresh_token)
    - [USER_PROFILE](#user_profile)
    - [WORK_EXPERIENCE](#work_experience)
    - [TECHNOLOGY](#technology)
    - [USER_PROFILE_TECHNOLOGY](#user_profile_technology)
    - [NOTIFICATION_SETTINGS](#notification_settings)
  - [Project & Recruitment](#project--recruitment)
    - [PROJECT](#project)
    - [PROJECT_ROLE](#project_role)
    - [INVITATION](#invitation)
    - [ROLE_APPLICATION](#role_application)
  - [Collaboration & Communication](#collaboration--communication)
    - [GROUP_CHAT](#group_chat)
    - [GROUP_CHAT_MEMBER](#group_chat_member)
    - [MESSAGE](#message)
  - [Enum Types](#enum-types)
  - [Relationships Summary](#relationships-summary)
  - [Security Notes](#security-notes)

---

## Overview

The **InnoSync** database schema supports user authentication, detailed talent profiles and project recruitment workflows.
It is designed for flexibility, data integrity, and extensibility to scale with platform features.

---

## Core Tables

### USERS

Stores basic authentication data for all users. (Table name: `users`)

| Field       | Type      | Description                 |
|-------------|-----------|-----------------------------|
| id          | bigint    | Primary key                 |
| email       | string    | Unique email                |
| full_name   | string    | Full name                   |
| password    | string    | Hashed password             |
| created_at  | timestamp | Registration timestamp      |
| last_login  | timestamp | Last login time (nullable)  |

---

### REFRESH_TOKEN

Stores JWT refresh tokens linked to users and devices.

| Field       | Type      | Description                      |
|-------------|-----------|----------------------------------|
| id          | bigint    | Primary key                      |
| user_id     | bigint    | Foreign key → users(id)          |
| token       | string    | Refresh token                    |
| expiry_date | timestamp | Expiration date                  |
| device_id   | string    | Device identifier (nullable)     |

---

### USER_PROFILE

Stores detailed user profile info beyond authentication.

| Field         | Type      | Description                                |
|---------------|-----------|--------------------------------------------|
| id            | bigint    | Primary key                                |
| user_id       | bigint    | Foreign key → users(id), unique            |
| telegram      | string    | Telegram handle (optional)                 |
| github        | string    | GitHub profile (optional)                  |
| bio           | string    | Short biography (optional)                 |
| position      | string    | Current position (optional)                |
| education     | education_enum | Enum: NO_DEGREE, BACHELOR, MASTER, PHD  |
| expertise     | string    | Expertise areas (free text or tags)        |
| expertise_level | expertise_level_enum | Enum: ENTRY, JUNIOR, MID, SENIOR, RESEARCHER |
| resume_url    | string    | URL to uploaded resume (optional)          |
| avatar_url    | string    | URL to profile picture (optional)          |

---

### WORK_EXPERIENCE

Records professional experiences for user profiles.

| Field           | Type      | Description                        |
|-----------------|-----------|------------------------------------|
| id              | bigint    | Primary key                        |
| user_profile_id | bigint    | Foreign key → user_profile(id)     |
| start_date      | date      | Employment start date               |
| end_date        | date      | Employment end date (nullable)      |
| position        | string    | Job title                          |
| company         | string    | Name of company                    |
| description     | string    | Job description (optional)          |

---

### TECHNOLOGY

Catalog of technologies used for tagging user skills and project needs.

| Field   | Type   | Description      |
|---------|--------|------------------|
| id      | bigint | Primary key      |
| name    | string | Technology name  |
| category| string | Technology category (optional) |

---

### USER_PROFILE_TECHNOLOGY

Join table for many-to-many relationship between user profiles and technologies.

| Field           | Type   | Description                    |
|-----------------|--------|--------------------------------|
| user_profile_id | bigint | Foreign key → user_profile(id) |
| technology_id   | bigint | Foreign key → technology(id)   |

---

### NOTIFICATION_SETTINGS

Stores notification preferences for each user (new for mobile app).

| Field                | Type    | Description                                 |
|----------------------|---------|---------------------------------------------|
| user_id              | bigint  | Primary key, Foreign key → users(id)        |
| new_message          | boolean | Notify on new message (default: true)       |
| application_update   | boolean | Notify on application update (default: true)|
| invitation_received  | boolean | Notify on invitation received (default: true)|
| project_update       | boolean | Notify on project update (default: true)    |

---

## Project & Recruitment

### PROJECT

Stores projects created by recruiters.

| Field           | Type      | Description                      |
|-----------------|-----------|----------------------------------|
| id              | bigint    | Primary key                      |
| recruiter_id    | bigint    | Foreign key → users(id)          |
| title           | string    | Project title                    |
| description     | string    | Project description (optional)   |
| team_size       | integer   | Number of required team members  |
| created_at      | timestamp | Project creation time            |
| is_active       | boolean   | Soft delete flag (default: true) |

---

### PROJECT_ROLE

Defines required roles and expectations per project.

| Field           | Type      | Description                     |
|-----------------|-----------|---------------------------------|
| id              | bigint    | Primary key                     |
| project_id      | bigint    | Foreign key → project(id)       |
| role_name       | string    | Role name (e.g., Frontend Dev)  |
| expertise_level | expertise_level_enum | Required expertise level |
| technologies    | string    | Technologies required (CSV)     |
| is_open         | boolean   | Role is open for applications (default: true) |

---

### INVITATION

Tracks project invitations sent from recruiters to recruitees.

| Field           | Type      | Description                      |
|-----------------|-----------|----------------------------------|
| id              | bigint    | Primary key                      |
| project_role_id | bigint    | Foreign key → project_role(id)   |
| user_id         | bigint    | Foreign key → users(id)          |
| status          | invitation_status_enum | INVITED, ACCEPTED, DECLINED, REVOKED |
| sent_at         | timestamp | Invitation send time              |
| responded_at    | timestamp | Time user responded (nullable)    |
| message         | string    | Personal invitation message (optional) |

---

### ROLE_APPLICATION

Tracks user applications to project roles. (Table name: `role_application`)

| Field           | Type      | Description                      |
|-----------------|-----------|----------------------------------|
| id              | bigint    | Primary key                      |
| user_id         | bigint    | Foreign key → users(id)          |
| project_role_id | bigint    | Foreign key → project_role(id)   |
| status          | application_status_enum | PENDING, UNDER_REVIEW, ACCEPTED, REJECTED, WITHDRAWN |
| applied_at      | timestamp | Application submission time      |
| updated_at      | timestamp | Last status update time          |
| cover_letter    | string    | Application details (optional)   |

---

## Collaboration & Communication

### GROUP_CHAT

Stores group chat info for project teams.

| Field       | Type      | Description                 |
|-------------|-----------|-----------------------------|
| id          | bigint    | Primary key                 |
| project_id  | bigint    | Foreign key → project(id)   |
| created_at  | timestamp | Group chat creation time    |
| chat_name   | string    | Custom chat name (optional) |
| is_active   | boolean   | Chat is active (default: true) |

---

### GROUP_CHAT_MEMBER

Maps users to group chats.

| Field         | Type      | Description                 |
|---------------|-----------|-----------------------------|
| group_chat_id | bigint    | Foreign key → group_chat(id) |
| user_id       | bigint    | Foreign key → users(id)      |
| joined_at     | timestamp | Join time (default: now)     |

---

### MESSAGE

Stores messages for group chats.

| Field        | Type      | Description                       |
|--------------|-----------|-----------------------------------|
| id           | bigint    | Primary key                       |
| sender_id    | bigint    | Foreign key → users(id)           |
| group_chat_id| bigint    | Foreign key → group_chat(id)      |
| content      | string    | Message text                      |
| sent_at      | timestamp | Message sent time                 |
| is_read      | boolean   | Read receipt (default: false)     |

---

## Enum Types

- **education_enum**: `NO_DEGREE`, `BACHELOR`, `MASTER`, `PHD`
- **expertise_level_enum**: `ENTRY`, `JUNIOR`, `MID`, `SENIOR`, `RESEARCHER`
- **invitation_status_enum**: `INVITED`, `ACCEPTED`, `DECLINED`, `REVOKED`
- **application_status_enum**: `PENDING`, `UNDER_REVIEW`, `ACCEPTED`, `REJECTED`, `WITHDRAWN`

---

## Relationships Summary

| Relationship                         | Cardinality          |
|---------------------------------------|----------------------|
| users → refresh_token                 | One-to-Many          |
| users → user_profile                  | One-to-One           |
| user_profile → work_experience        | One-to-Many          |
| user_profile → technology             | Many-to-Many         |
| users (Recruiter) → project           | One-to-Many          |
| project → project_role                | One-to-Many          |
| project_role → invitation             | One-to-Many          |
| users → role_application              | One-to-Many          |
| project_role → role_application       | One-to-Many          |
| project → group_chat                  | One-to-One           |
| group_chat → group_chat_member        | One-to-Many          |
| users → message                       | One-to-Many          |
| group_chat → message                  | One-to-Many          |
| users → notification_settings         | One-to-One           |

---

## Security Notes

- Passwords are stored hashed using secure algorithms (e.g., bcrypt).
- Tokens in REFRESH_TOKEN must be securely generated and stored.
- Enum sets are locked down at the database level for data integrity.
- Role-based access control is implemented at the backend layer.

---
