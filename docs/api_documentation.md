# Task-Manager

[Postman API Documentation for Task Manager API](https://documenter.getpostman.com/view/36830207/2sA3s7hnzM)

Welcome to the Task Manager API documentation. This API allows users to manage tasks efficiently by providing endpoints to create, read, update, and delete tasks. The API is designed to handle task-related operations such as setting task priorities, tracking due dates, and updating task statuses. Built with Go and utilizing the Gin framework, the Task Manager API ensures high performance and scalability for your task management needs.

## Get All Tasks

`GET {{URL}}/api/v1/tasks`

### API Request Description

This API endpoint makes an HTTP GET request to retrieve a list of tasks.

#### Request Parameters

- No request parameters are required for this endpoint.

#### Query Parameters

- No query parameters are required for this endpoint.

### API Response

The API returns a JSON object containing a list of tasks. Each task is represented by an object with the following properties:

- `id` (number): The unique identifier of the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `status` (string): The status of the task.
- `priority` (string): The priority of the task.
- `due_date` (string): The due date of the task.
- `created_at` (string): The timestamp when the task was created.
- `updated_at` (string): The timestamp when the task was last updated.

### Example

#### Request

```json
GET {{URL}}/api/v1/tasks
```

### Response

```json
{
[
  {
    "id": 1,
    "title": "Complete project report",
    "description": "Finish the annual project report by the end of the week",
    "status": "in progress",
    "priority": "high",
    "due_date": "2024-08-20T17:00:00Z",
    "created_at": "2024-08-10T09:00:00Z",
    "updated_at": "2024-08-14T12:00:00Z"
  },
  {
    "id": 2,
    "title": "Prepare presentation",
    "description": "Prepare slides for the upcoming presentation",
    "status": "pending",
    "priority": "medium",
    "due_date": "2024-08-22T10:00:00Z",
    "created_at": "2024-08-12T11:00:00Z",
    "updated_at": "2024-08-15T14:00:00Z"
  }
]
}
```

## Create Task

`{{URL}}api/v1/tasks/`

### Create Task

This endpoint allows the user to create a new task.

#### Request Body

- `id` (number, required): The unique identifier for the task.
- `title` (string, required): The title of the task.
- `description` (string, required): A brief description of the task.
- `status` (string, required): The status of the task (e.g., completed, pending, in progress).
- `priority` (string, required): The priority level of the task (e.g., high, medium, low).
- `due_date` (string, required): The due date for the task in the format "YYYY-MM-DDTHH:MM:SSZ".
- `created_at` (string, required): The date and time when the task was created in the format "YYYY-MM-DDTHH:MM:SSZ".
- `updated_at` (string, required): The date and time when the task was last updated in the format "YYYY-MM-DDTHH:MM:SSZ".

### Example

#### Request

```json
POST {{URL}}/api/v1/tasks

{
  "title": "Write unit tests",
  "description": "Write unit tests for the new features implemented",
  "status": "pending",
  "priority": "high",
  "due_date": "2024-08-20T17:00:00Z"
}
```

### Response

```json
{
  "type": "object",
  "properties": {
    "id": { "type": "number" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "status": { "type": "string" },
    "priority": { "type": "string" },
    "due_date": { "type": "string" },
    "created_at": { "type": "string" },
    "updated_at": { "type": "string" }
  }
}
```

### Body (raw json)

```json
{
  "id": 432,
  "title": "Complete project report",
  "description": "Finish the annual project report by the end of the week",
  "status": "completed",
  "priority": "high",
  "due_date": "0001-01-01T00:00:00Z",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```

##Get Single Task
`{{URL}}api/v1/tasks/:id`

### Example

#### Request

```json

GET {{URL}}/api/v1/tasks/:id

```

#### Response

- Status: 200 OK
- Content-Type: application/json
- Body:

  ```json
  {
    "id": 0,
    "title": "",
    "description": "",
    "status": "",
    "priority": "",
    "due_date": "",
    "created_at": "",
    "updated_at": ""
  }
  ```

This endpoint returns the details of a task identified by the provided ID, including its title, description, status, priority, due date, creation timestamp, and last update timestamp.

#### JSON Schema

```json
{
  "type": "object",
  "properties": {
    "id": {
      "type": "number"
    },
    "title": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "status": {
      "type": "string"
    },
    "priority": {
      "type": "string"
    },
    "due_date": {
      "type": "string"
    },
    "created_at": {
      "type": "string"
    },
    "updated_at": {
      "type": "string"
    }
  }
}
```

### Path Variables

`id              {{randomTen}}`

## Update Single Task

`{{URL}}api/v1/tasks/:id`

### Update Task

This endpoint is used to update a specific task by its ID.

### Example

#### Request

```json
PUT {{URL}}/api/v1/tasks/1

{
"title": "Update project report",
"description": "Update the project report with the latest data",
"status": "in progress",
"priority": "high",
"due_date": "2024-08-22T17:00:00Z"
}

```

#### Response

```json
{
  "id": 1,
  "title": "Update project report",
  "description": "Update the project report with the latest data",
  "status": "in progress",
  "priority": "high",
  "due_date": "2024-08-22T17:00:00Z",
  "created_at": "2024-08-10T09:00:00Z",
  "updated_at": "2024-08-15T09:00:00Z"
}
```

The response is a JSON object with the following schema:

```json
{
  "type": "object",
  "properties": {
    "id": { "type": "number" },
    "title": { "type": "string" },
    "description": { "type": "string" },
    "status": { "type": "string" },
    "priority": { "type": "string" },
    "due_date": { "type": "string" },
    "created_at": { "type": "string" },
    "updated_at": { "type": "string" }
  }
}
```

### Path Variables

`id                 {{randomTen}}`

## Delete Single Task

`{{URL}}api/v1/tasks/:id`

### Delete Task

This endpoint is used to delete a specific task by providing the task ID in the URL.

### Example

#### Request

```json
DELETE {{URL}}/api/v1/tasks/1
```

#### Response

```json
{
  "success": true
}
```

### Path Variables

`id                 {{randomTen}}`
