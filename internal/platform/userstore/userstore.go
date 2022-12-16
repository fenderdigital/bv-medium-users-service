package userstore

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/branebrvl/medium-users-service/internal"
)

type UserRow struct {
	PK    string `json:"pk"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersDB interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

type UsersClient struct {
	API UsersDB
}

func (c *UsersClient) Create(id, name, email string) (*internal.User, error) {
	in := c.BuildUserRow(id, name, email)
	out, err := c.API.PutItem(in)
	if err != nil {
		return nil, fmt.Errorf("PutItem failed: %w", err)
	}

	return c.MapUser(*out), nil
}

func (c *UsersClient) BuildUserRow(id, name, email string) *dynamodb.PutItemInput {
	return nil
}

func (c *UsersClient) MapUser(item dynamodb.PutItemOutput) *internal.User {
	return nil
}
