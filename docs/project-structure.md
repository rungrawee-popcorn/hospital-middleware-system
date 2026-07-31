# Project Structure

## Overview

Hospital Middleware System is developed using Go and Gin Framework following a layered architecture to improve maintainability, scalability, and code readability.

```
hospital-middleware-system
│
├── cmd
│   └── server
│       └── main.go
│
├── configs
│   ├── config.go
│   └── database.go
│
├── internal
│   ├── controllers
│   ├── services
│   ├── repositories
│   ├── models
│   ├── middleware
│   ├── hospital
│   ├── routes
│   └── utils
│
├── tests
│
├── docs
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Folder Description

### cmd/server

Application entry point.

Responsible for

- loading configuration
- connecting database
- initializing dependencies
- starting Gin server

---

### configs

Application configuration.

Contains

- Environment variables
- Database connection

---

### internal/controllers

Handle HTTP requests and responses.

Responsibilities

- Validate request
- Call service layer
- Return JSON response

---

### internal/services

Business logic.

Responsibilities

- Staff authentication
- Patient searching
- Hospital validation
- Hospital A API integration

---

### internal/repositories

Database access layer.

Responsibilities

- Query PostgreSQL
- CRUD operations

---

### internal/models

Database models.

Contains

- Hospital
- Staff
- Patient

---

### internal/middleware

Authentication middleware.

Current middleware

- JWT Authentication

---

### internal/routes

Application routes.

Responsible for registering APIs.

---

### internal/hospital

Hospital A API client.

Responsibilities

- Call external API
- Parse response
- Return standard model

---

### internal/utils

Common utilities.

Contains

- JWT Token Generator

---

### tests

Unit tests.

Contains

- Staff API tests
- Patient API tests

---

### docs

Project documentation.

Contains

- API Specification
- ER Diagram
- Project Structure

---

## Architecture

Client

↓

Gin Router

↓

Controller

↓

Service

↓

Repository

↓

PostgreSQL

↓

Hospital A API (when patient is not found locally)
