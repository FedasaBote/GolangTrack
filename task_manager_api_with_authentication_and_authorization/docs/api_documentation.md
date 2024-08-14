# Task Management API

## General Description
The Task Management API allows users to manage tasks by creating, reading, updating, and deleting tasks. Each task has attributes such as ID, title, description, due date, and status. This API provides endpoints for performing CRUD operations on tasks.

### Base URL

[http://localhost:8080](http://localhost:8080)

## Instructions
### Seting up MongoDB
**1. Install MongoDB**
- Follow the instructions on MongoDB website to install MongoDB - [MongoDB installation guide](https://www.mongodb.com/docs/manual/installation/)

**2. Start MongoDB Server**
- Start mongodb server using `mongod` command

**3. Install Project Dependencies**
- Install Project dependencies using command:
```sh
go get .
```

### Setting up Environment Variables
Ensure that you have a `.env` file in the root of your project with the following variables:
```sh
JWT_SECRET=your_jwt_secret_key
MONGO_URI=mongodb://localhost:27017
```

## Project Structure
```sh
    task_manager/
    ├── main.go
    ├── controllers/
    │   └── task_controller.go
    ├── models/
    │   └── task.go
    ├── data/
    │   └── task_service.go
    ├── router/
    │   └── router.go
    ├── docs/
    │   └── api_documentation.md
    └── go.mod
```

- **main.go:** Entry point of the application.
- **controllers/controller.go:** Handles incoming HTTP requests and invokes the appropriate service methods for both tasks and user authentication.
- **models/task.go:** Defines the Task struct.
- **models/user.go:** Defines the User struct.
- **data/task_service.go:** Contains business logic and data manipulation functions.
- **data/user_service.go:** Contains business logic and data manipulation functions for users, including hashing passwords.

- **middleware/auth_middleware.go:** Implements middleware to validate JWT tokens for authentication and authorization.
- **router/router.go:** Sets up the routes and initializes the Gin router and defines the routing configuration for the API.
- **docs/api_documentation.md:** Contains API documentation and other related documentation.
- **go.mod:** Defines the module and its dependencies.

## Authentication & Authorization
## JWT Authentication
The API uses JWT (JSON Web Tokens) for authenticating users. A valid JWT must be included in the `Authorization` header of the request to access protected routes.

## User Roles
- **Admin:** Has full access to all endpoints, including creating, updating, deleting tasks, and promoting users to admin.
- **User:** Can only access endpoints to retrieve all tasks or retrieve a task by its ID.
## Authorization Middleware
The `AuthMiddleWare` checks if the incoming request contains a valid JWT. The `RoleMiddleware` ensures that only users with the appropriate roles can access certain endpoints.

## Creating the First Admin User
If the database is empty, the first user to register will automatically be assigned the role of `admin`.

## Endpoints

### 1. User Registration
**URL:** `/register`  
**Method:** `POST`  
**Description:** Registers a new user. The first registered user will be an admin, and subsequent users will be regular users.

**Request:**  
- **Body:**
```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response:**
- **Status:** 201 Created
- **Body:**
```json
{
    "id": "60b488a18f682f0c1fc9cd35",
    "email": "user@example.com",
    "role": "admin"
}
```

### 2. User Login
**URL:** `/login`  
**Method:** `POST`  
**Description:** Authenticates a user and returns a JWT.

**Request:**  
- **Body:**
```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "message": "User logged in successfully!",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.    eyJlbWFpbCI6ImRhZ21hd2l0ZW5zYXlAZ21haWwuY29tIiwiaWQiOiI2NmI1ZmMyZWJjYzExYzRjODhmZDk1NmUiLCJyb2xlIjoiYWRtaW4ifQ.sTWX9W0a2zsKx8YQ9Fv_TzxZq0-0lv7FXQ1IFaXyaE0"
}
```

### 3. Get All Tasks
**URL:** `/tasks`  
**Method:** `GET`  
**Description:** Retrieves a list of all tasks.  
**Request:**  
- **Authorization:** Bearer `<JWT token>`

**Response:**
- **Status:** 200 OK
- **Body:**
```json
[
    {
        "id": "66b488a18f682f0c1fc9cd35",
        "title": "Task 1",
        "description": "description for task1",
        "due_date": "2024-08-08T07:23:10Z",
        "status": "pending"
    },
    {
        "id": "66b4ac6403193664f9a0f721",
        "title": "Task 2",
        "description": "description for task2",
        "due_date": "2024-09-14T07:23:10Z",
        "status": "done"
    }
]
```

### 4. Get a Task by ID
**URL:** `/tasks/:id`  
**Method:** `GET`  
**Description:**  Retrieves a task by its ID.
**Request:** 
- **Authorization:** Bearer `<JWT token>`
- **Parameters:** `id (string)` - The ID of the task.

**Response:**
- **Status:** 200 OK
- **Body:**
```json
 {
        "id": "66b488a18f682f0c1fc9cd35",
        "title": "Task 1",
        "description": "description for task1",
        "due_date": "2024-08-08T07:23:10Z",
        "status": "pending"
},
```
### 5. Create a New Task
**URL:** `/tasks/`  
**Method:** `POST`  
**Description:**  Creates a new task.
**Request:** 
- **Authorization:** Bearer `<JWT token>` (Admin role required)t
- **Body:**
```json
{
    "title": "Task1 ",
    "description": "description for task1",
    "due_date": "2024-08-08T07:23:10Z",
    "status": "pending"
}
```
**Response:**
- **Status:** 201 Created
- **Body:**
```json
{
    "id": "66b488a18f682f0c1fc9cd35",
    "title": "Task 1",
    "description": "description for task1",
    "due_date": "2024-08-10T12:00:00Z",
    "status": "pending"
}
```
### 6. Update a Task
**URL:** `/tasks/:id`  
**Method:** `PUT`  
**Description:**  Updates an existing task.
**Request:** 
- **Authorization:** Bearer `<JWT token>` (Admin role required)
- **Parameters:**
    - `id (string)` - The ID of the task.
- **Body:**
```json
{
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-12T12:00:00Z",
    "status": "completed"
}

```
**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "id": "66b488a18f682f0c1fc9cd35",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-12T12:00:00Z",
    "status": "completed"
}
```

### 7. Delete a Task
**URL:** `/tasks/:id`  
**Method:** `DELETE`  
**Description:**  Deletes a task by its ID.
**Request:** 
- **Authorization:** Bearer `<JWT token>` (Admin role required)
- **Parameters:**
    - `id (string)` - The ID of the task.
**Response:**
- **Status:** 204 No Content

### 8. Promote User to Admin
**URL:** `/users/promote/:id`  
**Method:** `PUT`  
**Description:**  Promotes a user to an admin role.
**Request:** 
- **Authorization:** Bearer `<JWT token>` (Admin role required)
- **Parameters:**
    - `id (string)` - The ID of the user to be promoted.
**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "message": "User promoted successfully!"
}
```