package auth

// Request/response DTOs for auth handlers.

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     int    `json:"role"`
	RoleName string `json:"role_name"`
}
