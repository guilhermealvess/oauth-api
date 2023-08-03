package entity

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length  = 40
)

type Client struct {
	ID                    uuid.UUID `validate:"required"`
	Name                  string    `validate:"required"`
	Description           string
	Slug                  string    `validate:"required"`
	IsActive              bool      `validate:"required"`
	ClientKey             string    `validate:"required"`
	HashClientSecret      string    `validate:"required"`
	RoundHashClientSecret int       `validate:"min=1"`
	SaltHashClientSecret  string    `validate:"required"`
	LastLogin             time.Time `validate:"required"`
	CreatedBy             string    `validate:"required"`
	CreatedAt             time.Time `validate:"required"`
	UpdatedAt             time.Time `validate:"required"`
	Permissions           []*Permission
}

func NewClient(name, desc, author string, permissions []*Permission) (*Client, string, error) {
	now := time.Now().UTC()
	round := rand.Intn(20) + 1
	client := Client{
		ID:                    uuid.New(),
		Name:                  name,
		Description:           desc,
		Slug:                  name,
		IsActive:              true,
		RoundHashClientSecret: round,
		ClientKey:             "",
		HashClientSecret:      "",
		SaltHashClientSecret:  "",
		CreatedBy:             author,
		CreatedAt:             now,
		UpdatedAt:             now,
		Permissions:           permissions,
	}

	client.generateSecrets()
	secret := randomString(length)
	client.HashClientSecret = client.HashSecret(secret)

	return &client, secret, client.Ok()
}

func (c *Client) Ok() error {
	return validator.New().Struct(c)
}

func (c *Client) generateSecrets() {
	c.ClientKey = randomString(length)
	c.SaltHashClientSecret = randomString(length)
}

func (c *Client) HashSecret(secret string) string {
	for i := 0; i < c.RoundHashClientSecret; i++ {
		secret = secret + c.SaltHashClientSecret

		hash := sha1.New()
		hash.Write([]byte(secret))
		hashSum := hash.Sum(nil)
		secret = hex.EncodeToString(hashSum)
	}
	return secret
}

func randomString(l int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}