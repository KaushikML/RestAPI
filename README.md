# Event REST API

A simple API built with Go and Gin providing user authentication and CRUD operations for events. Data is stored in SQLite.

## Quick Start

1. Fetch dependencies:

   ```bash
   go mod tidy
   ```
2. Run the server:

   ```bash
   go run main.go
   ```

The server will listen on `http://localhost:8080`.

## API Overview

| Method | Endpoint               | Description                  |
| ------ | ---------------------- | ---------------------------- |
| POST   | `/signup`              | Register a new user          |
| POST   | `/login`               | Authenticate and receive JWT |
| GET    | `/events`              | List all events              |
| GET    | `/events/:id`          | Get event by ID              |
| POST   | `/events`              | Create a new event *(auth)*  |
| PUT    | `/events/:id`          | Update an event *(auth)*     |
| DELETE | `/events/:id`          | Delete an event *(auth)*     |
| POST   | `/events/:id/register` | Register for event *(auth)*  |
| DELETE | `/events/:id/register` | Cancel registration *(auth)* |

> Every endpoint returns JSON. For protected routes, send the token returned from `/login` in the `Authorization` header.

## Testing with REST Client(Preferred) 

Inside `api-test/` you will find `.http` files for each endpoint. Install the **REST Client** extension in VS Code, open a `.http` file and click **Send Request** to try the API without crafting curl commands manually.

## Sample `curl` Usage

### `curl.exe` Raw Commands (Single-line)(DO NOT USE POWERSHELL , THERE IS SOME ISSUE WITH IT)

These commands work in cmd.exe without line continuations:

**1. Signup**

```bash
curl.exe -X POST http://localhost:8080/signup -H "Content-Type: application/json" --data "{\"email\":\"user3@example.com\",\"password\":\"secret\"}"
```

**2. Login** (to view full JSON response)

```bash
curl.exe -v -X POST http://localhost:8080/login -H "Content-Type: application/json" --data "{\"email\":\"user3@example.com\",\"password\":\"secret\"}"
```

> Copy the value of the `token` field from the response.

**3. Create Event**

```bash
curl.exe -X POST http://localhost:8080/events -H "Authorization: Bearer <YOUR_TOKEN_HERE>" -H "Content-Type: application/json" --data "{\"name\":\"Party\",\"description\":\"Fun\",\"location\":\"Town\",\"dateTime\":\"2025-01-01T15:30:00Z\"}"
```





## Simple Frontend

A minimal frontend is located in `frontend/`. Serve it with:

```bash
cd frontend
python3 -m http.server 8081
```

Open `http://localhost:8081` in your browser to interact with the API via forms.

A minimal frontend is located in `frontend/`. Serve it with:

```bash
cd frontend
python3 -m http.server 8081
```

Open `http://localhost:8081` in your browser to interact with the API via forms.

A minimal frontend is located in `frontend/`. Serve it with:

```bash
cd frontend
python3 -m http.server 8081
```

Open `http://localhost:8081` in your browser to interact with the API via forms.
