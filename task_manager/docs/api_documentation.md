# Task Management API

## General Description

The Task Management API allows users to manage tasks by creating, reading, updating, and deleting tasks. Each task has attributes such as ID, title, description, due date, and status.

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

### 2. Get a Task by ID

**URL:** `/tasks/:id`  
**Method:** `GET`  
**Description:** Retrieves a task by its ID.
**Request:**

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

### 3. Create a New Task

**URL:** `/tasks/`  
**Method:** `POST`  
**Description:** Creates a new task.
**Request:**

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

### 4. Update a Task

**URL:** `/tasks/:id`  
**Method:** `PUT`  
**Description:** Updates an existing task.
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
  "id": "66b488a18f682f0c1fc9cd35",
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-08-12T12:00:00Z",
  "status": "completed"
}
```

### 5. Delete a Task

**URL:** `/tasks/:id`  
**Method:** `DELETE`  
**Description:** Deletes a task by its ID.
**Request:**

- **Parameters:** - `id (string)` - The ID of the task.
  **Response:**
- **Status:** 204 No Content
