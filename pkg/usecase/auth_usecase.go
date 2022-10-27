package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"order-app/middleware"

	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"
	"order-app/pkg/repository"
	"os"
	"time"
)

type AuthUsecase interface {
	Login(request dto.LoginRequest) (dto.LoginResponse, error)
	Register(request dto.Register) (mysql.User, error)
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

func (u *authUsecase) Login(request dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		user mysql.User
		response dto.LoginResponse
		err  error
	)

	user, err = u.userRepository.Get(request.Username)
	if err != nil {
		if(err.Error() == "record not found"){
			err = errors.New("User not found")
			return response, err
		}
		return response, err
	}

	if user.Password != request.Password{
		err = errors.New("Wrong password")
		return response, err
	}

	response.Token, err = u.generateToken(user)
	if err != nil{
		return response, err
	}

	return response, nil
}

func (u *authUsecase) Register(request dto.Register) (mysql.User, error) {
	var (
		user mysql.User
		err  error
	)

	user, err = u.userRepository.Get(request.Username)
	if err != nil {
		if(err.Error() != "record not found"){
			return user, err
		}
	}

	if user.Id != "" {
		err = errors.New("Username has been used")
		return user, err
	}

	user.Username = request.Username
	user.Password = request.Password
	user.Name = request.Name

	err = u.userRepository.Insert(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *authUsecase) generateToken(user mysql.User) (string, error) {
	claims := middleware.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APPLICATION_NAME"),
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserId: user.Id,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, err
}
