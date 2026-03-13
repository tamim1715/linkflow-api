# LinkFlow API

Production-ready passwordless authentication service built with Go, Echo, and MongoDB.

LinkFlow demonstrates secure Magic Link authentication, JWT-based authorization, abuse prevention, clean layered
architecture, and containerized deployment.

---

## Table of Contents

- Overview
- Features
- Architecture
- Authentication Flow
- API Endpoints
- Security Considerations
- Project Structure
- Configuration
- Development Setup
- Docker
- Design Decisions
- Production Readiness
- Future Improvements
- Author

---

## Overview

LinkFlow implements a secure authentication flow using expiring magic links instead of passwords.

Authentication Flow:

1. User requests login link
2. Server generates a time-limited token
3. Token is delivered via deep link
4. User verifies the token
5. Server issues JWT
6. JWT is used to access protected endpoints

The project emphasizes clarity, maintainability, and production-minded engineering practices.

---

## Features

- Passwordless authentication (Magic Link)
- JWT-based stateless authorization
- Protected feedback endpoint
- MongoDB persistence
- TTL-based automatic token expiration
- One active token per user (replay prevention)
- Per-IP rate limiting on authentication endpoint
- Environment-based configuration
- Multi-stage Docker build

---

## Architecture

The system follows a layered architecture:

```
Transport Layer (Echo Handlers)
        ↓
Application Layer (Services)
        ↓
Domain Layer (Entities & Interfaces)
        ↓
Infrastructure Layer (MongoDB, Email, Slack)
```

### Architectural Principles

- Handlers contain no business logic
- Services encapsulate domain rules
- Repositories abstract data persistence
- Infrastructure details are isolated
- Dependencies are injected manually (constructor pattern)
- Configuration is environment-driven

---

## Authentication Flow

### 1. Request Magic Link

```
POST /api/auth/request-link
```

Request:

```json
{
    "email": "shahadathhossain447@gmail.com"
}
```

Behavior:

- Finds or auto-creates user
- Invalidates previous unused tokens
- Generates new expiring token
- Stores token in MongoDB
- Sends deep link (mocked email sender)

---

### 2. Verify Magic Link

```
GET /api/auth/verify?token=xxxx
```

Response:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzI4ODgzODksInVzZXJJZCI6ImViNDAzYjdmLTA3ZWYtNGE5Yi04MmZjLTFiNGVhY2I0YTQ1YyJ9.xB9ONWgclw0VvssCN3If4WqmQVdlIOlZ04JpW2NEEG4"
}
```

Behavior:

- Validates token existence
- Checks expiration
- Prevents reuse
- Marks token as used
- Generates JWT

---

### 3. Protected Feedback Endpoint

```
POST /api/feedback
Authorization: Bearer <jwt>
```

Request:

```json
{
    "message": "Great system!, but feedback typing issue"
}
```

---

## Security Considerations

- Token expiration handled via MongoDB TTL index
- Only one active token per user
- Token replay protection
- Per-IP rate limiting on authentication endpoint
- JWT-based stateless authentication
- Secrets managed via environment variables
- No password storage

---

## Project Structure

```
cmd/
└── server/                 # Application entry point

internal/
    app/                    # Server & route registration
    config/                 # Environment configuration
    constants/              # Centralized constants
    domain/                 # Core business entities
    handler/                # HTTP transport layer
    middleware/             # JWT & rate limiting
    repository/
        mongo/              # MongoDB implementations
    service/                # Business logic
    database/               # Database initialization
```

This structure enforces separation of concerns and long-term maintainability.

---

## Configuration

Create `.env` file for development:

```
APP_ENV=development
SERVER_PORT=8080
JWT_EXPIRES_HOURS=168
JWT_SECRET=super-secret-key
MONGODB_DATABASE=linkflow
MONGODB_URI=mongodb://localhost:27017
VERIFY_TOKEN_URI=myapp://api/auth/verify

```

In production, use system environment variables instead of `.env`.

---

## Development Setup

Clone repository:

```
git clone https://github.com/tamim1715/linkflow-api.git
cd linkflow-api
```

Run MongoDB locally:

```
mongod
```

Start server:

```
go run ./cmd/server
```

Server runs at:

```
http://localhost:8080
```

---

## Docker

Build image:

```
docker build -t tamim447/linkflow:v1.1.1  .
```

Run container:

```
docker run -p 8080:8080 \
-e APP_ENV=production \
-e SERVER_PORT=8080 \
-e MONGODB_URI=mongodb://host.docker.internal:27017 \
-e MONGODB_DATABASE=linkflow \
-e JWT_SECRET=strong-secret \
linkflow
```

The Dockerfile uses a multi-stage build to produce a minimal runtime image.

---

## Design Decisions

Why Magic Link?

- Eliminates password storage risks
- Simplifies user onboarding

Why JWT?

- Stateless authentication
- Scalable and infrastructure-friendly

Why MongoDB TTL Index?

- Automatic token cleanup
- No background jobs required

Why Manual Dependency Injection?

- Explicit wiring
- No framework lock-in
- Better testability

Why Rate Limiting Only on Auth?

- Authentication endpoints are most vulnerable to abuse
- Protects system resources

---

## Production Readiness

This project demonstrates:

- Layered architecture
- Abuse prevention
- Token lifecycle management
- Environment-driven configuration
- Infrastructure isolation
- Containerized deployment

---

## Future Improvements

- Redis-based distributed rate limiting
- Refresh token rotation
- Real email provider integration
- Structured logging (zap or zerolog)
- Observability and metrics
- Kubernetes deployment configuration

---

## Author

Shahadath Hossain Tamim  
Software Engineer
