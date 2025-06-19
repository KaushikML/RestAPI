# REST API with Go, Gin and SQLite

This project implements a small REST API for managing events. It includes user authentication and CRUD endpoints for events, plus a simple browser frontend.

## Database

The server uses SQLite and automatically creates `api.db` when started. Tables for users, events and registrations are created if they do not exist.

## Running the server

```bash
go run main.go
```

The server listens on `http://localhost:8080`.

## API Endpoints

All endpoints accept and return JSON. Authentication is handled via a JWT token returned from `/login` and must be sent in the `Authorization` header for protected routes.

### `POST /signup`
Create a new user.

Request body:
```json
{ "email": "user@example.com", "password": "secret" }
```
Response example:
```json
{ "message": "User created successfully" }
```

### `POST /login`
Authenticate a user and receive a token.

Request body is the same as `/signup`.
Response example:
```json
{ "message": "Login successful!", "token": "<jwt>" }
```

### `GET /events`
Return all events.

### `GET /events/:id`
Return a single event by id.

### `POST /events` *(authenticated)*
Create a new event.

Request body:
```json
{ "name": "Party", "description": "Fun", "location": "Town", "dateTime": "2025-01-01T15:30:00Z" }
```

### `PUT /events/:id` *(authenticated)*
Update an existing event.

### `DELETE /events/:id` *(authenticated)*
Delete an event.

### `POST /events/:id/register` *(authenticated)*
Register the current user for an event.

### `DELETE /events/:id/register` *(authenticated)*
Cancel a registration.

## Example `curl` commands

```bash
# create a user
curl -X POST http://localhost:8080/signup -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"secret"}'

# login
TOKEN=$(curl -s -X POST http://localhost:8080/login -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"secret"}' | jq -r '.token')

# create an event
curl -X POST http://localhost:8080/events -H "Authorization: $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"name":"Party","description":"Fun","location":"Town","dateTime":"2025-01-01T15:30:00Z"}'
```

## Frontend

A small frontend is located in the `frontend/` directory. Open `index.html` in a browser or serve the folder with a static file server:

```bash
cd frontend
python3 -m http.server 8081
```

Then navigate to `http://localhost:8081` and use the forms to interact with the API.

