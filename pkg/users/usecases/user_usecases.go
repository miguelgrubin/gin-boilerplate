// Package usecases provides use cases for petshop module.
package usecases

import (
	"log"

	sd "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/repositories"
)

type UserUseCases struct {
	ur repositories.UserRepository
	js services.JWTService
	hs services.HashService
}

var _ UserUseCasesInterface = &UserUseCases{}

type UserCreatorParams struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
}

type UserUpdatersParams struct {
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	Phone     *string
	Status    *string
}

type UserUseCasesInterface interface {
	Creator(UserCreatorParams) (domain.User, error)
	Showher(string) (domain.User, error)
	Updater(string, UserUpdatersParams) (domain.User, error)
	Deleter(string) error
	LoggerIn(string, string) (string, string, error)
	RefreshToken(string) (string, string, error)
	LoggerOut(string) error
}

func NewUserUseCases(pr repositories.UserRepository, js services.JWTService, hs services.HashService) UserUseCases {
	return UserUseCases{ur: pr, js: js, hs: hs}
}

func (p *UserUseCases) Creator(params UserCreatorParams) (domain.User, error) {
	user := domain.NewUser(domain.CreateUserParams(params))
	hashPwd, err := p.hs.Hash(params.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.PasswordHash = hashPwd
	err = p.ur.Save(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (p *UserUseCases) Showher(username string) (domain.User, error) {
	user, err := p.ur.FindOneByUsername(username)
	if err != nil {
		return domain.User{}, &domain.UsernameNotFound{Username: username}
	}
	return *user, nil
}

func (p *UserUseCases) Updater(username string, payload UserUpdatersParams) (domain.User, error) {
	user, err := p.ur.FindOneByUsername(username)
	if err != nil {
		return domain.User{}, &domain.UsernameNotFound{Username: username}
	}
	user.Update(domain.UpdateUserParams(payload))
	if payload.Password != nil {
		hashPwd, err := p.hs.Hash(*payload.Password)
		if err != nil {
			return domain.User{}, err
		}
		user.PasswordHash = hashPwd
	}
	err = p.ur.Save(*user)
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

func (p *UserUseCases) Deleter(username string) error {
	user, err := p.ur.FindOneByUsername(username)
	if err != nil {
		return &domain.UsernameNotFound{Username: username}
	}
	return p.ur.Delete(user.ID)
}

func (p *UserUseCases) LoggerIn(username string, password string) (string, string, error) {
	user, err := p.ur.FindOneByUsername(username)
	if err != nil {
		log.Println("username not found")
		return "", "", &domain.InvalidLogin{}
	}

	if !p.hs.Compare(user.PasswordHash, password) {
		log.Println("invalid password")
		return "", "", &domain.InvalidLogin{}
	}

	jwt, refreshToken, err := p.js.GenerateTokens(user.ID, user.Role)

	return jwt, refreshToken, err
}

func (p *UserUseCases) RefreshToken(refreshToken string) (string, string, error) {
	data, err := p.js.DecodeToken(refreshToken)
	if err != nil {
		return "", "", &sd.InvalidRefreshToken{}
	}

	user, err := p.ur.FindOne(data.UserID)
	if err != nil {
		return "", "", &sd.InvalidRefreshToken{}
	}

	newToken, newRefreshToken, err := p.js.RefreshToken(refreshToken, user.ID, user.Role)
	if err != nil {
		return "", "", &sd.InvalidRefreshToken{}
	}

	return newToken, newRefreshToken, err
}

func (p *UserUseCases) LoggerOut(token string) error {
	err := p.js.InvalidateToken(token)
	return err
}
