package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
)

const baseURL = "https://gateway-api.circle.com/v1"

const (
	balancesPath = "/balances"
	transferPath = "/transfer"
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

func (c *Client) Balances(ctx context.Context, address common.Address, domains []enum.CircleDomain) (map[enum.CircleDomain]*big.Int, error) {
	u := buildBalancesURL()
	req := newBalancesRequest(address, domains)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	out := balancesResponse{}
	if err := c.doRequest(ctx, u, http.MethodPost, bytes.NewReader(body), &out); err != nil {
		return nil, err
	}

	return out.Map(), nil
}

func (c *Client) CreateTransferAttestation(ctx context.Context, intents []models.BurnIntent) (TransferRequest, error) {
	u := buildTransferURL()
	body, err := json.Marshal(intents)
	if err != nil {
		return TransferRequest{}, fmt.Errorf("marshal request: %w", err)
	}

	out := TransferRequest{}
	if err := c.doRequest(ctx, u, http.MethodPost, bytes.NewReader(body), &out); err != nil {
		return TransferRequest{}, err
	}

	return out, nil
}

func (c *Client) doRequest(ctx context.Context, url string, method string, body io.Reader, out any) error {
	c.rate.Take()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
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

func buildTransferURL() string {
	return buildURL(transferPath)
}

func buildBalancesURL() string {
	return buildURL(balancesPath)
}

func buildURL(path string) string {
	return baseURL + path
}
