# Golang Backend API (Clean Architecture)

## 📌 Overview

This repository demonstrates a **production-grade Golang backend service** designed using **clean architecture principles**. The project is inspired by real-world enterprise systems (media / content platforms) and focuses on **scalability, maintainability, and clarity**.

> Note: This project is a **sanitized, public version** created for demonstration purposes only. No proprietary or confidential information is included.

---

## 🧠 Problem Statement

Modern backend systems require:

* Clear separation of concerns
* Scalable service layers
* Safe database migrations
* Easy testability and maintenance

This project showcases how to structure such a backend service using Golang.

---

## 🏗 Architecture

The project follows a **layered clean architecture**:

```
cmd/                → Application entry point
internal/
  handler/          → HTTP handlers (controllers)
  service/          → Business logic
  repository/       → Database access layer
  model/            → Domain models
migrations/         → SQL migrations
config/             → Configuration loading
pkg/                → Shared utilities
```

Flow:

```
HTTP Request → Handler → Service → Repository → Database
```

---

## 🛠 Tech Stack

* **Language:** Golang
* **Framework:** net/http / Gin
* **Database:** PostgreSQL
* **Migrations:** SQL-based migrations
* **Containerization:** Docker

---

## ✨ Key Features

* Clean architecture folder structure
* Role & permission–style domain modeling
* PostgreSQL schema with migrations
* Config-driven application setup
* Ready for extension with auth, middleware, and observability

---

## 🚀 Getting Started

### Prerequisites

* Go 1.21+
* PostgreSQL

### Run Locally

```bash
# clone repo
git clone https://github.com/<your-username>/golang-backend-api.git
cd golang-backend-api

# run application
go run cmd/api/main.go
```

---

## 📂 Example API (Sample)

```
GET /health
Response: 200 OK
```

---

## 🎯 Why This Project Exists

This repository is created to:

* Demonstrate **real-world Golang backend practices**
* Showcase **system design & code organization**
* Support my transition into **Senior / Lead Backend Engineer roles**

---

## 👤 Author

**Rohit Sharma**
Senior Backend Engineer
Node.js | Golang | MongoDB | AWS

---

## 📄 License

MIT
