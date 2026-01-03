package server

import "github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserCreateRequest struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UserUpdateRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Phone     *string `json:"phone"`
	Status    *string `json:"status"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func UserResponseFromDomain(u domain.User) UserResponse {
	return UserResponse{
		u.ID,
		u.Username,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		u.Status,
		u.Role,
		u.CreatedAt.String(),
		u.UpdatedAt.String(),
	}
}
