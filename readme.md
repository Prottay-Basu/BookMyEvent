# BookMyEvent

## Overview
BookMyEvent is a functional web application tailored for managing event bookings efficiently. It enables users to create, view, update, and delete events while providing seamless user authentication and secure interactions.

## Features
- **Event Management**:
  - Create, view, update, and delete events.
  - Store event details, including name, description, location, and datetime.
  - Associate events with specific users for better tracking.

- **User Authentication**:
  - Secure login and registration with JWT-based authentication.
  - Passwords are securely hashed using bcrypt to prevent unauthorized access.

- **Database Operations**:
  - SQLite integration for lightweight and efficient data storage.
  - Connection pooling and validation to ensure database stability and performance.

- **API Functionality**:
  - Comprehensive API endpoints to handle all CRUD operations for users and events.
  - Organized route definitions with middleware support for authentication.

- **Middleware Support**:
  - Authentication middleware to validate JWT tokens and ensure secure API access.
  - Modular middleware design for easy extension and maintenance.

- **Utility Functions**:
  - Password hashing and verification using bcrypt.
  - JWT token generation and validation for session management.

- **Testing Support**:
  - Predefined HTTP requests to validate the functionality of APIs.
  - Includes tests for event creation, user registration, and authentication.

- **Scalable Design**:
  - Modular structure with dedicated directories for models, routes, utilities, and middleware.
  - Easy to extend and maintain, supporting future feature additions.

## Project Structure
```
BookMyEvent/
├── api-test       # API test cases for validating endpoints
├── db             # Database configuration and initialization
├── go.mod         # Go module dependencies
├── go.sum         # Dependency checksums
├── main.go        # Entry point of the application
├── middlewares    # Middleware functions for authentication
├── models         # Database models for events and users
├── routes         # API route definitions
├── utils          # Utility functions for hashing and JWT
```

## Getting Started

### Prerequisites
- **Go**: Version 1.23 or higher.
- **SQLite**: Installed locally.

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd BookMyEvent
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application
1. Start the application:
   ```bash
   go run main.go
   ```
2. The server will be available at `http://localhost:8080`.

## Testing
- Test cases are located in the `api-test` directory.
- Run tests using:
  ```bash
  go test ./api-test/...
  ```

