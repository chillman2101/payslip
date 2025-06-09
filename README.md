# Payslip System Documentation

## Table of Contents
1. [How-To Guides](#how-to-guides)
2. [API Usage](#api-usage)
3. [Software Architecture](#software-architecture)

---

## How-To Guides

### 1. How to Run the Project

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-org/payslip-system.git
   cd payslip-system
   ```

2. **Set up `.env`**:
   ```env
    DB=postgresql://postgres:postgres@localhost:5432/db_payslip
    SERVER_PORT=8081
    REDIS_URL=redis://default@127.0.0.1:6379
    AUTH_KEY=asdasdasdasd
   ```

3. **Run migrations**:
   ```bash
   go run main.go migrate
   ```

4. **Start the server**:
   ```bash
   go run main.go
   ```

---

### 2. How to Run Tests

- **Unit tests**:
  ```bash
  go test ./tests/unit/... -short
  ```

- **Integration tests**:
  Ensure `db_payslip` DB exists and configured in `.env`, then:
  ```bash
  go test ./tests/integration/...
  ```

---

### 3. How to Create an Admin & Employee
Uncomment the "Seeder(db)" code to perform Seeder 100 employee and 2 admin accounts
```go
func NewDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := config.DB

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	Migration(db)
	// Seeder(db) 

	return db, nil
}
```

---

## API Usage

### Auth

#### POST `/api/login/admin`
- **Description**: Login as admin.
- **Request**:
```json
{
  "username": "admin",
  "password": "admin"
}
```
- **Response**:
```json
{
    "token": "<jwt-token>",
}
```

#### POST `/api/login/employee`
- **Description**: Login as employee.
- **Request**:
```json
{
  "username": "bmabe0",
  "password": "bmabe0"
}
```
- **Response**:
```json
{
  "token": "<jwt-token>",  
}
```

---

### Payslip Employee Summary

#### GET `/employee/payslip/generate/:id`
- id:
 - **Description**: Payroll ID.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
```json
{
    "data": {
        "employee_id": 85,
        "employee_name": "ascarsbrooke2c",
        "attendance": {
            "attendances": [
                {
                    "attendance_date": "2025-06-09T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-10T00:00:00Z",
                    "check_in_time": "2025-06-10 20:12:57",
                    "check_out_time": "2025-06-10 20:43:50"
                },
                {
                    "attendance_date": "2025-06-11T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-12T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-13T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-14T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-15T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-16T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-17T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-18T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-19T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-20T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-21T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-22T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-23T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                },
                {
                    "attendance_date": "2025-06-24T00:00:00Z",
                    "check_in_time": "2025-06-09 20:12:57",
                    "check_out_time": "2025-06-08 20:43:50"
                }
            ],
            "total_day_present": 21,
            "total_working_day": 21,
            "total": 2054000
        },
        "overtime": {
            "overtimes": [
                {
                    "overtime_date": "2025-06-08T00:00:00Z",
                    "overtime_hours": 2,
                    "start_time": "2025-06-08 20:55:03",
                    "end_time": "2025-06-08 22:55:03"
                }
            ],
            "total_overtime_hour": 2,
            "total_working_hour": 168,
            "total": 48904
        },
        "reimbursement": {
            "reimbursements": null,
            "total": 0
        },
        "salary": 2054000,
        "total_take_home_pay": 2102904
    },
    "message": "successfully get summary"
}
```

### Payslip Admin Summary

#### GET `/admin/payslip/generate/:id`
- id:
 - **Description**: Payroll ID.
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
```json
{
    "data": {
        "take_home_pay_employee": [
            {
                "employee_id": 85,
                "employee_name": "ascarsbrooke2c",
                "total_take_home_pay": 2102904
            }
        ],
        "total_take_home_pay_all_employee": 2102904
    },
    "message": "successfully get summary"
}
```

---

## Software Architecture

### Layered Architecture

```
┌──────────────────────┐
│      HTTP Handler    │ ← Gin Framework
└──────────────────────┘
           ↓
┌──────────────────────┐
│       Service        │ ← Business Logic
└──────────────────────┘
           ↓
┌──────────────────────┐
│     Repository       │ ← DB Queries (GORM)
└──────────────────────┘
           ↓
┌──────────────────────┐
│ Database (Postgres)  │
└──────────────────────┘
```

### Components

| Layer     | Responsibility                                    |
|-----------|---------------------------------------------------|
| Handler   | Route binding, request validation                 |
| Service   | Implements business logic                         |
| Repository| Interacts with the database via GORM              |
| Models    | Entity definitions                                |
| Middleware| JWT auth, request tracing, audit logging          |

---

### Security & Logging

- **JWT Auth** for employee login.
- **Password Hashing** using `bcrypt`.
- **Audit Logging** via middleware (logs `RequestID`, actor, endpoint, and timestamp).

---

### Observability

- Each request includes a `RequestID` in logs.
- Logs printed with actor (admin/employee), action, and duration.
- Middleware for:
  - Request tracing
  - Performance logging
  - Checking Active Payroll
  - Checking Auth
