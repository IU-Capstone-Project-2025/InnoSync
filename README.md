# InnoSync

## Table of Contents
- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Key User Roles](#key-user-roles)
- [FastSyncing](#fastsyncing-aiml-powered-auto-match)
- [Collaboration Tools](#collaboration-tools)
- [Similar Platforms](#similar-platforms)
- [Competitive Advantages](#what-sets-innosync-apart)
- [Setting up the project](#setting-up-the-project)

## Overview

**InnoSync** is a smart talent-matching and project collaboration platform tailored for university communities and innovation hubs. It connects _recruiters_ (project creators) with _recruitees_ (talent) by streamlining team formation through detailed profile creation, role-based filtering, and AI-powered fast matchmaking.

## Tech Stack

### Frontend (Asgat & Anvar)
- React
- Next.js
- TypeScript
- Figma (for UI/UX design)

### Backend (Yusuf & Egor)
- Java
- Spring Boot
- PostgreSQL

### DevOps (Aibek & Baha)
- TBA

### Machine Learning (Gurban)
- Python

## Key User Roles

### Recruitee (Talent/Applicant)

- Sign up and specify role as _Recruitee_
- Build a comprehensive profile including:
  - Desired positions (e.g., Frontend Developer, ML Engineer)
  - Tech stack and tools proficiency
  - Expertise level (Entry, Junior, Mid, Senior, Researcher)
  - LinkedIn and GitHub integration
- Browse open projects using advanced filters
- Apply directly to projects of interest
- Receive and manage invitations from recruiters

### Recruiter (Project Creator)

- Sign up and specify role as _Recruiter_
- Create a project with:
  - Title and description
  - Team size and required positions
  - Tech stack and technologies per role
  - Expertise level expectations (Entry to Researcher)
- Choose between:
  1. **Manual Search:** Browse and filter to discover talent
  2. **FastSyncing:** Use AI/ML to automatically match suitable team members

## FastSyncing (AI/ML-Powered Auto-Match)

- Recruiters opt-in to FastSyncing
- AI/ML models suggest optimal team members based on:
  - Role requirements
  - Past experience and availability
  - Skill compatibility
- Recruiters can:
  - Review full profiles
  - _Revoll_ any suggestion for a better fit
  - Send batch invites to proposed team
- Recruitees may accept or decline invitations

## Collaboration Tools

- Upon all invite acceptances:
  - Recruiter is notified to finalize the team
  - A group chat is created for all team members
  - One-on-one messaging is also enabled

## Similar Platforms

- LinkedIn
- Upwork
- Topta
- Freelancer
- Fiverr
- GitHub Projects

## What Sets InnoSync Apart

- **AI/ML-Powered Team Formation (FastSyncing):** Offers AI-powered team formation through FastSyncing, allowing recruiters to auto-match with suitable candidates based on roles, skills, and expertise levels.
- **Team-Centric Recruitment:** Focuses on building entire collaborative teams for projects, not just filling isolated job positions.
- **Integrated Onboarding and Communication:** Includes built-in group and individual chat features to streamline onboarding and collaboration without relying on external tools.
- Provides a fast, game-inspired user experience

## Setting up the project
Follow these steps to clone the **InnoSync** repository using either HTTPS or SSH.
### 1) Clone the Repository
#### **Option 1: HTTPS (Recommended for beginners)**
```bash
git clone https://github.com/IU-Capstone-Project-2025/InnoSync.git
```
#### **Option 2: SSH (Requires SSH key setup)**
```bash
git clone git@github.com:IU-Capstone-Project-2025/InnoSync.git
```
### 2) Navigate to the Project Directory
```bash
cd InnoSync
```
### 3) Create .env file in ```backend``` directory, contact us through telegram for credentials (@abdugafforzoda)
### 4)  Come back to root folder (```InnoSync```) and build and start containers
```bash
docker-compose up --build
```
### 5)  Make simple hello request
http://localhost:3000/

