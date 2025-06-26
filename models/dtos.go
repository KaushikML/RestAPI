package models

// ErrorResponse represents a generic API error.
// swagger:model
type ErrorResponse struct {
    // A human-readable description of the error
    Message string `json:"message" example:"invalid request"`
}

// Message is a simple success wrapper.
// swagger:model
type Message struct {
    Message string `json:"message" example:"success"`
}

// Token carries a signed JWT.
// swagger:model
type Token struct {
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// UserLogin is the request body for /login.
// swagger:model
type UserLogin struct {
    Email    string `json:"email"    example:"alice@test.com"`
    Password string `json:"password" example:"secret"`
}

// EventIn is the payload used when creating/updating an event.
// swagger:model
type EventIn struct {
    Name        string `json:"name"        example:"GinConf"`
    Description string `json:"description" example:"Annual Gin meetup"`
    Location    string `json:"location"    example:"Delhi"`
    DateTime    string `json:"dateTime"    example:"2025-07-01T10:00:00Z"`
}
