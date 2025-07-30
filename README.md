# Go Users REST API (No Framework, No ORM)

A lightweight RESTful API built using only Go's standard library. This application handles basic CRUD operations for users, storing data in-memory without using any frameworks, external libraries, or databases.

---

## Features

- No frameworks (Fiber, Gin, etc.)
- No ORMs (GORM, SQLBoiler, etc.)
- In-memory data store (slice of structs)
- Clean and minimal structure
- Full CRUD support via RESTful endpoints

---

## Endpoints

### ✅ Get All Users
**GET** `/users`  
**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
]

```
---

### 🔍 Get Single User

**GET** `/users?id=1`
**Response:**

* `200 OK` with user JSON if found
* `404 Not Found` if user does not exist

---

### ➕ Create User

**POST** `/users`
**Headers:**

```
Content-Type: application/json
```

**Body:**

```json
{
  "name": "Jane Smith",
  "email": "jane@example.com"
}
```

**Response:**

* `200 OK` with created user JSON

---

### ✏️ Update User

**PUT** `/users?id=1`
**Headers:**

```
Content-Type: application/json
```

**Body:**

```json
{
  "name": "Updated Name",
  "email": "updated@example.com"
}
```

**Response:**

* `200 OK` with updated user JSON
* `404 Not Found` if user does not exist

---

### ❌ Delete User

**DELETE** `/users?id=1`
**Response:**

* `204 No Content` on success
* `404 Not Found` if user does not exist

---

## Project Structure

```
.
├── main.go                 # Entry point, route bindings
├── models/
│   └── user.go             # User struct & in-memory logic
├── handlers/
│   └── user_handler.go     # HTTP handler functions
```

---

## Running the App

Start the server:

```bash
go run main.go
```

Server will be available at `http://localhost:3000`

---

## Tools Used

* Language: Go (Golang)
* Standard packages only:

  * `net/http`
  * `encoding/json`
  * `strconv`

---

## Example cURL Requests

**Create User**

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

**Get User**

```bash
curl http://localhost:3000/users?id=1
```

**Update User**

```bash
curl -X PUT http://localhost:3000/users?id=1 \
  -H "Content-Type: application/json" \
  -d '{"name":"New Name","email":"new@example.com"}'
```

**Delete User**

```bash
curl -X DELETE http://localhost:3000/users?id=1
```

---

## Warning

This project uses in-memory storage (slice). All data will be **reset when the server restarts**. This is intentional to keep things simple and dependency-free.

---
