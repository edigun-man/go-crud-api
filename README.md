# Go CRUD REST API (No Framework, No ORM)

A simple and clean REST API in Go — written with zero frameworks and no ORM — just standard library and raw SQL.

This project is ideal for learning how Go works under the hood when building web APIs.

---

## 🚀 Features

- 🧱 Built using only Go’s standard `net/http`
- 🧠 No third-party frameworks or routing libraries
- 💾 In-memory store for fast testing (optional: raw SQL database)
- 📦 JSON request/response handling with `encoding/json`
- 🧪 Easy to test using Postman or `curl`
- 💡 Beginner-friendly code structure

---

## 📁 Project Structure

```bash
.
├── main.go              # Entry point of the API server
├── models/
│   └── user.go          # User struct definition
├── handlers/
│   └── user_handler.go  # HTTP handlers for CRUD operations
├── db/
│   └── connection.go    # (Optional) DB connection setup using database/sql
