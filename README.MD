# Event Management API

A simple **RESTful API** built with **Go**, **Gin Gonic**, and **GORM** to manage events, with **PostgreSQL** as the database.  
This service includes endpoints to **create**, **list**, and **retrieve events by ID**, and integrates **Swagger** for API documentation.

---

## **Features**
- RESTful API with Go and Gin Gonic.
- PostgreSQL integration using GORM.
- Swagger (OpenAPI) documentation.
- Docker Compose setup for easy local environment setup.
- Functional API tests using Bruno collections.

---

## **Project Structure**
```
.
├── bruno_collection/       # Bruno API test collection
├── cmd/app/                # Application entrypoint
├── configs/                # Config and DB connection
├── controllers/            # HTTP handlers
├── docs/                   # Swagger generated files
├── models/                 # Data models
├── repository/             # Database layer
├── routes/                 # URL mapping
├── scripts/                # DB initialization scripts
├── services/               # Business logic
├── .env                    # Environment variables
├── docker-compose.yml      # Docker services
├── Makefile                # Build and run commands
└── README.md
```

---

## **Requirements**
- Go 1.23+
- Docker & Docker Compose
- [Bruno](https://www.usebruno.com/) (optional, for API testing)
- PostgreSQL (if running locally without Docker)

---

## **Setup & Run**

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/giojimen3z/lkt_test_interview.git
cd event-api
```

### 2️⃣ Environment Variables
Create a `.env` file based on `.env.example`:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=eventsdb
DB_SSLMODE=disable
```
---

## **Swagger Note**
The latest version of **swag** has an issue that adds two invalid fields in `docs/docs.go`:
```go
LeftDelim:  "{{",
RightDelim: "}}",
```
If you see an error like:
```
unknown field LeftDelim in struct literal of type "github.com/swaggo/swag".Spec
```
Simply **delete these two lines** from `docs/docs.go` and regenerate Swagger docs.

---

###  3️⃣ Generate Swagger Documentation
```bash
make swag
```
or:
```bash
swag init -g cmd/app/main.go -o docs
```

---


### 4️⃣ Run with Docker Compose
```bash
make compose-up
```
This will:
- Start PostgreSQL
- Run the API
- Apply database migrations

---

### 5️⃣ Access the API
- Health Check:
```bash
curl http://localhost:8080/ping
```
- Swagger UI:
```
http://localhost:8080/swagger/index.html
```

---

### 6️⃣ Stop Services
```bash
make compose-down
```

---

## **API Endpoints**
| Method | Endpoint         | Description          |
|--------|-----------------|----------------------|
| POST   | /events         | Create an event      |
| GET    | /events         | List all events      |
| GET    | /events/{id}    | Get event by ID      |
| GET    | /ping           | Health check         |

---

## **Testing API with Bruno**
We provide a **Bruno collection** under `bruno_collection/`.  
You can import it in Bruno or run manually:
```bash
bruno run bruno_collection/
```
*(Make sure Bruno CLI is installed)*

---

## **License**
MIT
