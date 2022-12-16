package pubsub

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Publisher interface {
	Publish(input *sns.PublishInput) (*sns.PublishOutput, error)
}

type Client struct {
	API      Publisher
	TopicArn *string
}

func (c *Client) Publish(msg string) error {
	if _, err := c.API.Publish(&sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: c.TopicArn,
	}); err != nil {
		return fmt.Errorf("publish failed: %w", err)
	}

	return nil
}
