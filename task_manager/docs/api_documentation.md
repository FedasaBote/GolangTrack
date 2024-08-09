# Task Management API

## General Description
The Task Management API allows users to manage tasks by creating, reading, updating, and deleting tasks. Each task has attributes such as ID, title, description, due date, and status. This API provides endpoints for performing CRUD operations on tasks.

### Base URL

[http://localhost:8080](http://localhost:8080)

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
- **controllers/task_controller.go:** Handles incoming HTTP requests and invokes the appropriate service methods.
- **models/:** Defines the data structures used in the application.
- **data/task_service.go:** Contains business logic and data manipulation functions.
- **router/router.go:** Sets up the routes and initializes the Gin router and defines the routing configuration for the API.
- **docs/api_documentation.md:** Contains API documentation and other related documentation.
- **go.mod:** Defines the module and its dependencies.

## Endpoints

### 1. Get All Tasks
**URL:** `/tasks`  
**Method:** `GET`  
**Description:** Retrieves a list of all tasks.  
**Request:**  
No parameters.

**Response:**
- **Status:** 200 OK
- **Body:**
```json
[
    {
        "id": "1",
        "title": "Task 1",
        "description": "Description of Task 1",
        "due_date": "2024-08-10T12:00:00Z",
        "status": "pending"
    },
    {
        "id": "2",
        "title": "Task 2",
        "description": "Description of Task 2",
        "due_date": "2024-08-11T12:00:00Z",
        "status": "completed"
    }
]
```

### 2. Get a Task by ID
**URL:** `/tasks/:id`  
**Method:** `GET`  
**Description:**  Retrieves a task by its ID.
**Request:** 
- **Parameters:** `id (string)` - The ID of the task.

**Response:**
- **Status:** 200 OK
- **Body:**
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-10T12:00:00Z",
    "status": "pending"
}
```
### 3. Create a New Task
**URL:** `/tasks/`  
**Method:** `POST`  
**Description:**  Creates a new task.
**Request:** 
- **Body:**
```json
{
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-10T12:00:00Z",
    "status": "pending"
}
```
**Response:**
- **Status:** 201 Created
- **Body:**
```json
{
    "id": "3",
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-10T12:00:00Z",
    "status": "pending"
}
```
### 4. Update a Task
**URL:** `/tasks/:id`  
**Method:** `PUT`  
**Description:**  Updates an existing task.
**Request:** 
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
    "id": "1",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-12T12:00:00Z",
    "status": "completed"
}
```

### 5. Delete a Task
**URL:** `/tasks/:id`  
**Method:** `DELETE`  
**Description:**  Deletes a task by its ID.
**Request:** 
- **Parameters:**
    - `id (string)` - The ID of the task.
**Response:**
- **Status:** 204 No Content

Link to Postman [API Documentation](https://documenter.getpostman.com/view/25805253/2sA3rzKCtU) 