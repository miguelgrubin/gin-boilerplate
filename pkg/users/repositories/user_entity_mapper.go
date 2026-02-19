package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
)

func UserEntityToDomain(ue UserEntity) domain.User {
	return domain.HydrateUser(
		ue.ID,
		ue.Username,
		ue.FirstName,
		ue.LastName,
		ue.Email,
		ue.Phone,
		ue.PasswordHash,
		ue.Status,
		ue.Role,
		ue.CreatedAt,
		ue.UpdatedAt,
		ue.DeletedAt,
	)
}

func UserEntityFromDomain(u domain.User) UserEntity {
	return UserEntity{
		ID:           u.ID,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Phone:        u.Phone,
		PasswordHash: u.PasswordHash,
		Status:       u.Status,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    u.DeletedAt,
	}
}
