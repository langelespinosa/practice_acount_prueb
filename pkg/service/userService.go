package service

import (
	"errors"
	"os"
	"practice_account/pkg/domain"
	"practice_account/pkg/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Encrypt(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func (s *UserService) CreateUser(user *domain.Users) error {
	return s.userRepo.Create(user)
}

func (s *UserService) GetUser(id uint) (*domain.Users, error) {
	return s.userRepo.FindById(id)
}

func (s *UserService) GetUserWithTrans(id uint) (*domain.Users, error) {
	return s.userRepo.FindByIdWithTrans(id)
}

func (s *UserService) ValidateEmail(local string) error {
	return s.userRepo.FindByEmail(local)
}

func (s *UserService) ValidateLogin(login string, pass string) (string, string, error) {

	userDb, err := s.userRepo.FindByLogin(login)
	if err != nil {
		return "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(pass)); err != nil {
		return "", "", err
	}

	keyLogin := os.Getenv("myQSQ7HH_SQeX7DrU84PPabTH9a3_oeFw84sWc4j0ys")

	tokenLogin := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userDb.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	tokenStringL, err := tokenLogin.SignedString([]byte(keyLogin))
	if err != nil {
		return "", "", err
	}

	keyAccess := os.Getenv("t9fpd4phF2owAVtNLo73hUjwBPsbZ2kUMfF75WeJQ0w")

	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userDb.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
	})

	tokenStringA, err := tokenAccess.SignedString([]byte(keyAccess))
	if err != nil {
		return "", "", err
	}

	return tokenStringL, tokenStringA, nil
}

func (s *UserService) ValidatePass(hashPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
}

func (s *UserService) GetAllUsers() ([]domain.Users, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetAllUsersWithTrans() ([]domain.Users, error) {
	return s.userRepo.FindAllWithTrans()
}

func (s *UserService) UpdateUser(id uint, user *domain.Users) error {
	return s.userRepo.Update(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) AuthRequired(ctx *fiber.Ctx) error {

	cookieAccess := ctx.Cookies("access")
	if cookieAccess == "" {

		cookieRefresh := ctx.Cookies("refresh")
		if cookieRefresh == "" {
			return errors.New("Unauthorized")
		}

		jwtKey := os.Getenv("myQSQ7HH_SQeX7DrU84PPabTH9a3_oeFw84sWc4j0ys")
		tokenRefresh, err := jwt.ParseWithClaims(cookieRefresh, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !tokenRefresh.Valid {
			return err
		}

		claims, ok := tokenRefresh.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))
		if !ok {
			return err
		}

		tokenAccessNew := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": userID,
			"exp":     jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		})

		jwtKeyA := os.Getenv("t9fpd4phF2owAVtNLo73hUjwBPsbZ2kUMfF75WeJQ0w")
		tokenStringA, err := tokenAccessNew.SignedString([]byte(jwtKeyA))
		if err != nil {
			return err
		}

		ctx.Cookie(&fiber.Cookie{
			Name:     "access",
			Value:    tokenStringA,
			Expires:  time.Now().Add(time.Minute * 30),
			HTTPOnly: true,
		})
	} else {

		jwtKeyA := os.Getenv("t9fpd4phF2owAVtNLo73hUjwBPsbZ2kUMfF75WeJQ0w")
		tokenAccess, err := jwt.ParseWithClaims(cookieAccess, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKeyA), nil
		})

		if !tokenAccess.Valid {
			return err
		}
	}

	return nil
}
