package cctp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"go.uber.org/ratelimit"
)

const baseURL = "https://iris-api.circle.com/v2"

const (
	feesPath     = "/burn/USDC/fees"
	messagesPath = "/messages"
)

var (
	ErrMessageNotFound    = errors.New("message not found")
	ErrStatusNotCompleted = errors.New("message status not completed")
)

type Client struct {
	rate   ratelimit.Limiter
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		rate:   ratelimit.New(10),
		client: &http.Client{},
	}
}

func (c *Client) Fees(ctx context.Context, srcDomain, dstDomain string) (Fees, error) {
	u := buildFeesURL(srcDomain, dstDomain)

	out := Fees{}
	if err := c.doRequest(ctx, u, http.MethodGet, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) MessageAndAttestation(ctx context.Context, srcDomain, txHash string) (string, string, error) {
	u := buildMessageAndAttestationURL(srcDomain, txHash)

	out := messagesResponse{}
	if err := c.doRequest(ctx, u, http.MethodGet, &out); err != nil {
		return "", "", fmt.Errorf("message request: %w", err)
	}

	if len(out.Messages) == 0 {
		return "", "", ErrMessageNotFound
	}

	if out.Messages[0].Status != "complete" {
		return "", "", ErrStatusNotCompleted
	}

	return out.Messages[0].Message, out.Messages[0].Attestation, nil
}

func (c *Client) doRequest(ctx context.Context, url string, method string, out any) error {
	c.rate.Take()

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("failed request: %d", resp.StatusCode)
	}
}

func buildFeesURL(srcDomain, dstDomain string) string {
	p := path.Join(feesPath, srcDomain, dstDomain)
	return buildURL(p)
}

func buildMessageAndAttestationURL(srcDomain, txHash string) string {
	p := path.Join(messagesPath, srcDomain)
	params := url.Values{}
	params.Set("transactionHash", txHash)
	return buildURL(p + "?" + params.Encode())
}

func buildURL(path string) string {
	return baseURL + path
}
