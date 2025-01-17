package sdk

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	ListModels() ([]ProviderModels, error)
	GenerateContent(provider, model, prompt string) (GenerateResponse, error)
}

type ClientImpl struct {
	baseURL    string
	httpClient *resty.Client
}

type ProviderModels struct {
	Provider string        `json:"provider"`
	Models   []interface{} `json:"models"`
}

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ResponseTokens struct {
	Role    string `json:"role"`
	Model   string `json:"model"`
	Content string `json:"content"`
}

type GenerateResponse struct {
	Provider string         `json:"provider"`
	Response ResponseTokens `json:"response"`
}

func NewClient(baseURL string) Client {
	return &ClientImpl{
		baseURL:    baseURL,
		httpClient: resty.New(),
	}
}

func (c *ClientImpl) ListModels() ([]ProviderModels, error) {
	resp, err := c.httpClient.R().
		SetResult([]ProviderModels{}).
		Get(fmt.Sprintf("%s/llms", c.baseURL))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to list models: %s", resp.Status())
	}

	return *resp.Result().(*[]ProviderModels), nil
}

func (c *ClientImpl) GenerateContent(provider, model, prompt string) (GenerateResponse, error) {
	request := GenerateRequest{
		Model:  model,
		Prompt: prompt,
	}

	resp, err := c.httpClient.R().
		SetBody(request).
		SetResult(GenerateResponse{}).
		Post(fmt.Sprintf("%s/llms/%s/generate", c.baseURL, provider))
	if err != nil {
		return GenerateResponse{}, err
	}

	if resp.StatusCode() != http.StatusOK {
		return GenerateResponse{}, fmt.Errorf("failed to generate content: %s", resp.Status())
	}

	return *resp.Result().(*GenerateResponse), nil
}
