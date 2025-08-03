# Go Event CRUD API

A RESTful API for event management built with Go, featuring user authentication, event creation/management, and attendee registration functionality.

## Features

- **User Authentication**: JWT-based authentication with registration and login
- **Event Management**: Create, read, update, and delete events
- **Attendee Management**: Register/unregister attendees for events
- **Database Migrations**: Automated database schema management
- **API Documentation**: Swagger/OpenAPI documentation
- **Secure**: Password hashing with bcrypt and JWT authentication

## Tech Stack

- **Language**: Go 1.24.5
- **Web Framework**: Gin
- **Database**: SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Migrations**: golang-migrate
- **Documentation**: Swagger/OpenAPI with swaggo
- **Environment Management**: godotenv

## Project Structure

```
├── cmd/
│   ├── api/              # API server application
│   │   ├── main.go       # Application entry point
│   │   ├── auth.go       # Authentication handlers
│   │   ├── events.go     # Event CRUD handlers
│   │   ├── routes.go     # Route definitions
│   │   ├── server.go     # HTTP server setup
│   │   └── middleware.go # Authentication middleware
│   └── migrate/          # Database migration tool
│       ├── main.go
│       └── migrations/   # SQL migration files
├── internal/
│   ├── database/         # Database models and operations
│   │   ├── users.go
│   │   ├── events.go
│   │   ├── attendees.go
│   │   └── modals.go
│   └── env/              # Environment variable utilities
├── docs/                 # Generated Swagger documentation
└── data.db              # SQLite database file
```

## Getting Started

### Prerequisites

- Go 1.24.5 or later
- SQLite3
- Air (optional, for live reloading during development)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-event-crud
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables** (optional)
   Create a `.env` file in the root directory:
   ```env
   PORT=6969
   JWT_SECRET=your-secret-key-here
   ```

4. **Install Air for live reloading** (optional but recommended)
   ```bash
   go install github.com/air-verse/air@latest
   ```

5. **Run database migrations**
   ```bash
   go run cmd/migrate/main.go up
   ```

6. **Start the server**
   
   With Air (recommended for development):
   ```bash
   air
   ```
   
   Or without Air:
   ```bash
   go run cmd/api/*.go
   ```

The API will be available at `http://localhost:6969`

## API Documentation

Once the server is running, you can access the Swagger documentation at:
- **Swagger UI**: `http://localhost:6969/swagger/index.html`

## API Endpoints

### Authentication
- `POST /api/v1/register` - Register a new user
- `POST /api/v1/login` - Login user

### Events (Requires Authentication)
- `GET /api/v1/events` - Get all events
- `GET /api/v1/events/:id` - Get event by ID
- `POST /api/v1/events` - Create new event
- `PUT /api/v1/events/:id` - Update event
- `DELETE /api/v1/events/:id` - Delete event

### Attendees (Requires Authentication)
- `POST /api/v1/events/:id/register` - Register for an event
- `DELETE /api/v1/events/:id/register` - Unregister from an event

## Database Schema

### Users Table
- `id` (Primary Key)
- `email` (Unique)
- `name`
- `password` (Hashed)

### Events Table
- `id` (Primary Key)
- `owner_id` (Foreign Key to Users)
- `name`
- `description`
- `date`
- `location`

### Attendees Table
- `id` (Primary Key)
- `user_id` (Foreign Key to Users)
- `event_id` (Foreign Key to Events)

## Usage Examples

### Register a new user
```bash
curl -X POST http://localhost:6969/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login
```bash
curl -X POST http://localhost:6969/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create an event (with JWT token)
```bash
curl -X POST http://localhost:6969/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Tech Conference 2025",
    "description": "Annual technology conference",
    "date": "2025-09-15T09:00:00Z",
    "location": "Convention Center"
  }'
```

## Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 6969)
- `JWT_SECRET`: Secret key for JWT token signing (default: "random-secret")

## Database Migrations

To manage database schema changes:

```bash
# Apply migrations
go run cmd/migrate/main.go up

# Rollback migrations
go run cmd/migrate/main.go down
```

## Development

### Adding New Migrations

1. Create new migration files in `cmd/migrate/migrations/`:
   - `000xxx_migration_name.up.sql`
   - `000xxx_migration_name.down.sql`

2. Run the migration:
   ```bash
   go run cmd/migrate/main.go up
   ```

### Regenerating Swagger Documentation

```bash
swag init -g cmd/api/main.go
```
