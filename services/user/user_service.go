package user

import (
	"net/http"
	"os"
	"time"

	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	"github.com/BarryDaBee/go_restaurant_management/internal/models"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/golang-jwt/jwt/v4"
)

type userService struct {
	Storage storage.Storage
	Secret  string
}

func NewUserService(storage storage.Storage, secret string) UserService {
	if secret == "" {
		secret = os.Getenv("SECRET")
	}
	return &userService{
		Storage: storage,
		Secret:  secret,
	}
}

type UserService interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (string, error)
	GetUserByID(id uint) (*models.User, error)
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (u *userService) Register(user *models.User) (*models.User, error) {
	if u.Storage.CheckExists(&models.User{}, "email = ?", user.Email) {
		return nil, utility.NewAppError("user with this email already exists", http.StatusBadRequest)
	}

	if err := user.HashPassword(); err != nil {
		return nil, utility.NewAppError("failed to hash password", http.StatusInternalServerError)
	}

	err := u.Storage.Create(user)
	if err != nil {
		return nil, utility.NewAppError("failed to create user", http.StatusBadRequest)
	}

	user.Password = ""
	return user, nil
}

func (u *userService) Login(email, password string) (string, error) {
	var user models.User

	err := u.Storage.FindOne(&user, "email = ?", email)
	if err != nil {
		return "", utility.NewAppError("invalid email or password", http.StatusUnauthorized)
	}

	if !user.CheckPassword(password) {
		return "", utility.NewAppError("invalid email or password", http.StatusUnauthorized)
	}

	token, err := u.generateJWT(user.ID)
	if err != nil {
		return "", utility.NewAppError("failed to generate token", http.StatusInternalServerError)
	}

	return token, nil
}

func (u *userService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := u.Storage.FindByID(&user, id)
	if err != nil {
		return nil, utility.NewAppError("user not found", http.StatusNotFound)
	}
	user.Password = ""
	return &user, nil
}

func (u *userService) generateJWT(userID uint) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.Secret))
}
