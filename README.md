# LinkFlow

Magic Link Authentication + JWT + Feedback API  
Built with Go, Echo, MongoDB.

---

## Features

- Magic Link Authentication
- JWT-based authorization
- Protected feedback endpoint
- MongoDB storage
- Clean Architecture
- Docker support
- Environment-based configuration

---

## Tech Stack

- Go 1.22
- Echo Framework
- MongoDB
- JWT
- Docker

---

## Project Structure

internal<br>
├── app/<br>
├── config/<br>
├── domain/<br>
├── handler/<br>
├── middleware/<br>
├── repository/<br>
│   └── mongo/<br>
├── service/<br>
└── database/

cmd/<br>
└── server/

---

## Setup (Development)

### 1. Clone repository

```bash
git clone https://github.com/tamim1715/linkflow-api.git
cd linkflow
```

### 1. Docker command
docker build -t tamim447/linkflow:v1.0.1 .
docker push tamim447/linkflow:v1.0.1
