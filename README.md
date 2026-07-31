# Hospital Middleware System

Hospital Middleware System is a backend service developed for searching and managing patient information from Hospital Information Systems (HIS).

The system provides hospital staff authentication, JWT-based authorization, patient search functionality, hospital access control, and integration with external Hospital API services.

---

# Tech Stack

## Backend

- Go
- Gin Framework
- GORM ORM

## Database

- PostgreSQL

## Authentication

- JWT Authentication
- Bcrypt Password Hashing

## Containerization

- Docker
- Docker Compose

## Web Server

- Nginx

## Testing

- Go Testing
- Testify

## API Testing

- Postman

## Development Tools

- VS Code
- Git
- GitHub

---

# System Architecture

Client

↓

Nginx

↓

Gin API Server

↓

Controller

↓

Service

↓

Repository

↓

PostgreSQL Database

Patient Search Flow:

Request

↓

JWT Middleware

↓

Patient Controller

↓

Patient Service

↓

Search PostgreSQL

↓

Found:
Return Patient Data

Not Found:
Check National ID / Passport ID

↓

Hospital A API

↓

Map Response

↓

Save Patient Data

↓

Return Response

---

# Features

## Hospital Staff

- Create staff account
- Staff login
- JWT token generation
- Password hashing using bcrypt

## Patient Search

Search patients using:

- National ID
- Passport ID
- First Name
- Middle Name
- Last Name
- Date of Birth
- Phone Number
- Email

## Hospital Access Control

- Each staff member belongs to one hospital.
- Staff can only search patients belonging to their assigned hospital.
- Hospital validation is performed using hospital_id from JWT claims.

## External Hospital API

- Connect with Hospital A API.
- Map external patient response.
- Store patient information into PostgreSQL.

---

Documentation:

- docs/project-structure.md
- docs/api-spec.md
- docs/er-diagram.md

---

# Database Design

The system contains three main entities:

## Hospitals

Stores hospital information.

Columns:

- id
- name
- created_at
- updated_at

## Staffs

Stores hospital staff accounts.

Columns:

- id
- username
- password
- hospital_id
- created_at
- updated_at

Relationship:

Hospital (1) ---- (Many) Staff

## Patients

Stores patient information.

Columns:

- id
- hospital_id
- patient_hn
- national_id
- passport_id
- first_name_th
- middle_name_th
- last_name_th
- first_name_en
- middle_name_en
- last_name_en
- date_of_birth
- phone_number
- email
- gender
- created_at
- updated_at

Relationship:

Hospital (1) ---- (Many) Patient

---

# Environment Setup

## Requirements

Install:

- Go
- Docker Desktop
- Git
- Postman

Create environment file:

.env

Example:

APP_PORT=8080

APP_ENV=development

DB_HOST=postgres

DB_PORT=5432

DB_USER=postgres

DB_PASSWORD=postgres

DB_NAME=hospital_db

JWT_SECRET=your_secret_key

HOSPITAL_A_API_URL=https://hospital-a.api.co.th/patient/search

---

# Run with Docker

Build and start services:

docker compose up --build

Services:

| Service    | Port |
| ---------- | ---- |
| Go API     | 8080 |
| PostgreSQL | 5432 |
| Nginx      | 80   |

Stop services:

docker compose down

---

# Run Without Docker

Install dependencies:

go mod download

Run application:

go run ./cmd/server

---

# API Endpoints

## Create Staff

POST /staff/create

Create a new hospital staff account.

## Staff Login

POST /staff/login

Authenticate staff and generate JWT token.

## Search Patient

GET /api/patient/search

Authentication required.

Example:

GET /api/patient/search?national_id=1234567890123

Authorization Header:

Authorization: Bearer <jwt_token>

---

# Testing

Run all tests:

go test ./...

Test scenarios include:

- Staff creation success case
- Staff creation validation case
- Staff login success case
- Staff login failure case
- Patient search success case
- Unauthorized access case
- Hospital access validation case

---

# Security

Implemented security practices:

- JWT authentication
- Password hashing with bcrypt
- Request validation
- Environment variables for secrets
- Database foreign key relationships
- Hospital-level authorization

---

# Development Notes

- Followed layered architecture pattern.
- Controllers handle HTTP requests only.
- Business logic is separated into service layer.
- Database operations are handled by repository layer.
- Sensitive information is stored using environment variables.

---
