# Go Student App

Go Student App is a REST API for managing schools, students, and courses. It is built with Gin, Gorm, PostgreSQL, and Swagger documentation.

## Tech Stack

- Go 1.23+
- Gin
- Gorm
- PostgreSQL
- Swaggo Swagger UI

## Project Structure

```text
.
├── config/                 # Database connection
├── school/                 # School domain, application, infrastructure, HTTP handlers
├── student/                # Student domain, application, infrastructure, HTTP handlers
├── course/                 # Course domain, application, infrastructure, HTTP handlers
├── docs/                   # Generated Swagger files
├── main.go                 # Application entry point
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.23 or newer
- PostgreSQL running locally
- A database named `coursesdb`

The current database settings are defined in `config/database.go`:

```text
host: localhost
port: 5432
user: postgres
password: 1234
database: coursesdb
```

## Setup

Install dependencies:

```bash
go mod tidy
```

Set an API bearer token:

```bash
export API_AUTH_TOKEN="change-me"
```

Run the API:

```bash
go run .
```

The server starts on:

```text
http://localhost:8080
```

## Swagger

Swagger UI is available at:

```text
http://localhost:8080/swagger/index.html
```

Click **Authorize** in Swagger UI and enter:

```text
Bearer change-me
```

Regenerate Swagger docs after changing API annotations:

```bash
swag init -g main.go
```

If `swag` is not installed:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## API Routes

Base path:

```text
/api/v1
```

### Schools

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/school` | Create a school |
| `GET` | `/schools` | Get all schools |
| `GET` | `/schools/{school_id}` | Get a school by ID |
| `PUT` | `/schools/{school_id}` | Update a school |
| `DELETE` | `/schools/{school_id}` | Delete a school |

### Students

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/schools/{school_id}/student` | Create a student in a school |
| `GET` | `/schools/{school_id}/students` | Get students by school ID |
| `GET` | `/schools/{school_id}/students/{student_id}` | Get a student by ID |
| `PUT` | `/schools/{school_id}/students/{student_id}` | Update a student |
| `DELETE` | `/schools/{school_id}/students/{student_id}` | Delete a student |

### Courses

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/schools/{school_id}/students/{student_id}/course` | Create a course for a student |
| `GET` | `/schools/{school_id}/students/{student_id}/courses` | Get courses by student ID |
| `GET` | `/schools/{school_id}/students/{student_id}/courses/{course_id}` | Get a course by ID |
| `PUT` | `/schools/{school_id}/students/{student_id}/courses/{course_id}` | Update a course |
| `DELETE` | `/schools/{school_id}/students/{student_id}/courses/{course_id}` | Delete a course |

## Example Requests

Create a school:

```bash
curl -X POST http://localhost:8080/api/v1/school \
  -H "Authorization: Bearer change-me" \
  -H "Content-Type: application/json" \
  -d '{"name":"Central School"}'
```

Create a student:

```bash
curl -X POST http://localhost:8080/api/v1/schools/1/student \
  -H "Authorization: Bearer change-me" \
  -H "Content-Type: application/json" \
  -d '{"name":"Ali","class":"10"}'
```

Create a course:

```bash
curl -X POST http://localhost:8080/api/v1/schools/1/students/1/course \
  -H "Authorization: Bearer change-me" \
  -H "Content-Type: application/json" \
  -d '{"title":"Mathematics"}'
```

Requests without the bearer token return `401 Unauthorized`.

## Verification

Run all package checks:

```bash
go test ./...
go vet ./...
```
