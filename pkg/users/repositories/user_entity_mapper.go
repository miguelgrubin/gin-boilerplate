package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
)

func UserEntityToDomain(ue UserEntity) domain.User {
	user := domain.NewUser()
	user.ID = ue.ID
	user.Username = ue.Username
	user.FirstName = ue.FirstName
	user.LastName = ue.LastName
	user.Email = ue.Email
	user.PasswordHash = ue.PasswordHash
	user.Status = ue.Status
	user.Role = ue.Role
	user.CreatedAt = ue.CreatedAt
	user.UpdatedAt = ue.UpdatedAt
	user.DeletedAt = ue.DeletedAt
	return user
}

func UserEntityFromDomain(u domain.User) UserEntity {
	return UserEntity{
		ID:           u.ID,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Status:       u.Status,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    u.DeletedAt,
	}
}
