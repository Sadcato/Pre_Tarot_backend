package arkapi

import (
	"context"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

type Client struct {
	client *arkruntime.Client
}

func NewClient() *Client {
	return &Client{
		client: arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY")),
	}
}

func (c *Client) Chat(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	req := model.ChatCompletionRequest{
		Model: os.Getenv("ENDPOINT_ID"),
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(systemPrompt),
				},
			},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(userPrompt),
				},
			},
		},
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return *resp.Choices[0].Message.Content.StringValue, nil
}
