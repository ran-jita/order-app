package usecase

import (
	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"
	"order-app/pkg/repository"
)

type AuthUsecase interface {
	GetLogin(username string) (mysql.User, error)
	PostLogin(request dto.Login) (mysql.User, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(
	userRepository repository.UserRepository,
) *authUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

func (u *authUsecase) GetLogin(username string) (mysql.User, error) {
	var (
		user mysql.User
		err  error
	)

	user, err = u.userRepository.Get(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *authUsecase) PostLogin(request dto.Login) (mysql.User, error) {
	var (
		user mysql.User
		err  error
	)

	user.Username = request.Username
	user.Password = request.Password

	err = u.userRepository.Insert(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
