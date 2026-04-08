package greenapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"green/internal/config"
	"io"
	"net/http"
)

const (
	formatGetSettingsPath      = "/waInstance%s/getSettings/%s"
	formatGetStateInstancePath = "/waInstance%s/getStateInstance/%s"
	formatSendMessagePath      = "/waInstance%s/sendMessage/%s"
	formatSendFileByURLPath    = "/waInstance%s/sendFileByUrl/%s"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(cfg *config.GreenAPIConfig) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

func (c *Client) GetSettings(ctx context.Context, idInstance, apiToken string) (map[string]any, error) {
	path := fmt.Sprintf(formatGetSettingsPath, idInstance, apiToken)
	var result map[string]any
	if err := c.do(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	return result, nil
}

func (c *Client) GetStateInstance(ctx context.Context, idInstance, apiToken string) (map[string]any, error) {
	path := fmt.Sprintf(formatGetStateInstancePath, idInstance, apiToken)
	var result map[string]any
	if err := c.do(ctx, "GET", path, nil, &result); err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	return result, nil
}

func (c *Client) SendMessage(
	ctx context.Context,
	idInstance, apiToken string,
	req SendMessageRequest,
) (*SendMessageResponse, error) {
	path := fmt.Sprintf(formatSendMessagePath, idInstance, apiToken)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	var result SendMessageResponse
	if err = c.do(ctx, "POST", path, bytes.NewReader(body), &result); err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	return &result, nil
}

func (c *Client) SendFileByUrl(
	ctx context.Context,
	idInstance, apiToken string,
	req SendFileByUrlRequest,
) (*SendFileByUrlResponse, error) {
	path := fmt.Sprintf(formatSendFileByURLPath, idInstance, apiToken)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	var result SendFileByUrlResponse
	if err = c.do(ctx, "POST", path, bytes.NewReader(body), &result); err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	return &result, nil
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader, dst any) error {
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		raw, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(raw))
	}
	if dst != nil {
		if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
