package entity

import (
	"math/rand"
	"strings"
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
	name = removeSpecialCharacters(name)
	client := Client{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
		Slug:        slugFy(name),
		IsActive:    true,
		CreatedBy:   author,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: permissions,
	}

	client.setSlug()
	client.generateSecrets()
	secret := randomString(length)
	client.HashClientSecret = calculateSHA256(secret, client.SaltHashClientSecret, client.RoundHashClientSecret)

	return &client, secret, client.Ok()
}

func (c *Client) Ok() error {
	return validator.New().Struct(c)
}

func (c *Client) setSlug() {
	slug := strings.ToLower(c.Name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	c.Slug = slug
}

func (c *Client) generateSecrets() {
	c.ClientKey = randomString(length)
	c.SaltHashClientSecret = randomString(length)
	round := rand.Intn(20) + 1
	c.RoundHashClientSecret = round
}
