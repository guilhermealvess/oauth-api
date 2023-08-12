package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/guilhermealvess/oauth-api/config"
	"github.com/tee8z/nullable"
)

type User struct {
	ID                uuid.UUID `validate:"required"`
	TeamID            uuid.UUID `validate:"required"`
	Name              string    `validate:"required"`
	Occupation        string    `validate:"required"`
	Email             string    `validate:"required"`
	HashPassword      string    `validate:"required"`
	RoundHashPassword int       `validate:"required"`
	SaltHashPassword  string    `validate:"required"`
	LastLogin         nullable.Time
	CreatedAt         time.Time `validate:"required"`
	UpdatedAt         time.Time `validate:"required"`
}

func NewUser(teamID uuid.UUID, name, occupation, email, password string) (*User, error) {
	now := time.Now().UTC()
	rand.Seed(time.Now().UnixNano())
	u := User{
		ID:                uuid.New(),
		TeamID:            teamID,
		Name:              removeSpecialCharacters(name),
		Occupation:        occupation,
		Email:             email,
		CreatedAt:         now,
		UpdatedAt:         now,
		RoundHashPassword: rand.Intn(20) + 1,
		SaltHashPassword:  randomString(20),
	}
	u.HashPassword = calculateSHA256(password, u.SaltHashPassword, u.RoundHashPassword)

	return &u, u.Ok()
}

func (u *User) Ok() error {
	return validator.New().Struct(u)
}

func (u *User) buildTokenPayload() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"exp":     time.Now().UTC().Add(time.Hour).Unix(),
		"email":   u.Email,
		"user_id": u.ID.String(),
		"team_id": u.TeamID.String(),
	})

	return token.SignedString([]byte(config.Config.Secret))
}

func (u *User) SignIn(password string) (string, error) {
	passwordEncoded := calculateSHA256(password, u.SaltHashPassword, u.RoundHashPassword)
	if passwordEncoded != u.HashPassword {
		return "", fmt.Errorf("Forbidden, user %s not authorized", u.Name)
	}
	return u.buildTokenPayload()
}
