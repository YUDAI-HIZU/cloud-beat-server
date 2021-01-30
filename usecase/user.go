package usecase

import (
	"app/config"
	"app/domain/models"
	"app/domain/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

type UserUseCase interface {
	Create(db *gorm.DB, user *models.User) (*models.User, error)
	GetByID(db *gorm.DB, id int) (*models.User, error)
	SignIn(db *gorm.DB, email string, password string) (*models.User, string, error)
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: u,
	}
}

func (u userUseCase) Create(db *gorm.DB, user *models.User) (*models.User, error) {
	encryptedPassword, err := generateEncryptedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.EncryptedPassword = encryptedPassword
	user, err = u.userRepository.Create(db, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUseCase) GetByID(db *gorm.DB, id int) (*models.User, error) {
	user, err := u.userRepository.GetByID(db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUseCase) SignIn(db *gorm.DB, email string, password string) (*models.User, string, error) {
	user, err := u.userRepository.GetByEmail(db, email)
	if err != nil {
		return nil, "", err
	}
	if err := compareHashAndPassWord(user.EncryptedPassword, password); err != nil {
		return nil, "", err
	}
	token, err := generateToken(user)
	return user, token, nil
}

func generateEncryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashAndPassWord(encryptedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func generateToken(user *models.User) (string, error) {
	id, err := strconv.Atoi(user.ID)
	if err != nil {
		return "", err
	}
	claims := &jwt.MapClaims{
		"exp":    time.Now().Add(24 * 7 * time.Hour).Unix(),
		"userID": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	return tokenString, err
}
