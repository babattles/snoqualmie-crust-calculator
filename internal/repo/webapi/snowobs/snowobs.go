package snowobs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	dateFormat     = "200601021504"
)

type Config struct {
	BaseURL    string
	Token      string
	Source     string
	HTTPClient *http.Client
}

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	source     string
}

func New(cfg Config) *Client {
	baseURL := cfg.BaseURL
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		token:      cfg.Token,
		source:     cfg.Source,
	}
}

func (c *Client) GetStationData(ctx context.Context, stids []string, startDate, endDate time.Time) (*Response, error) {
	if len(stids) == 0 {
		return nil, fmt.Errorf("snowobs: at least one stid is required")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("snowobs: parse base url: %w", err)
	}

	q := u.Query()
	q.Set("token", c.token)
	q.Set("source", c.source)
	q.Set("stid", strings.Join(stids, ","))
	q.Set("start_date", startDate.Format(dateFormat))
	q.Set("end_date", endDate.Format(dateFormat))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("snowobs: build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("snowobs: do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("snowobs: unexpected status %d", resp.StatusCode)
	}

	var out Response
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("snowobs: decode response: %w", err)
	}

	return &out, nil
}
