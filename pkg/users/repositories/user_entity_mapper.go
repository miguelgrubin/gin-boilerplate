package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
)

func UserEntityToDomain(ue UserEntity) domain.User {
	return domain.User{
		ID:        ue.ID,
		Username:  ue.Username,
		FirstName: ue.FirstName,
		LastName:  ue.LastName,
		Email:     ue.Email,
		Status:    ue.Status,
		Role:      ue.Role,
		CreatedAt: ue.CreatedAt,
		UpdatedAt: ue.UpdatedAt,
		DeletedAt: ue.DeletedAt,
	}
}

func UserEntityFromDomain(u domain.User) UserEntity {
	return UserEntity{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Status:    u.Status,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
