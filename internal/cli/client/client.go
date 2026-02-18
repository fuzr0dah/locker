package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fuzr0dah/locker/internal/api"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) url(endpoint api.Endpoint) string {
	return c.baseURL + endpoint.Path
}

func (c *Client) CreateSecret(ctx context.Context, req *api.CreateSecretRequest) (*api.Secret, error) {
	url := c.url(api.Secrets.Create)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, api.Secrets.Create.Method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("decode error")
	}

	var secret api.Secret
	if err := json.NewDecoder(resp.Body).Decode(&secret); err != nil {
		return nil, err
	}
	return &secret, nil
}

func (c *Client) GetSecret(ctx context.Context, id string) (*api.Secret, error) {
	// TODO: multi-param support
	url := strings.Replace(c.url(api.Secrets.Get), "{id}", id, 1)

	httpReq, err := http.NewRequestWithContext(ctx, api.Secrets.Get.Method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, api.SecretNotFoundErr
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("decode error") // TODO c.decodeError(resp)
	}

	var secret api.Secret
	if err := json.NewDecoder(resp.Body).Decode(&secret); err != nil {
		return nil, err
	}
	return &secret, nil
}
