package platformclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"yangfan-ui/internal/config"
)

type Client struct {
	store      *config.Store
	httpClient *http.Client
}

func New(store *config.Store) *Client {
	return &Client{
		store: store,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Do(ctx context.Context, method string, path string, reqBody any, extraHeaders map[string]string) (http.Header, []byte, error) {
	if c.store == nil {
		return nil, nil, errors.New("config store 未初始化")
	}
	baseURL := c.store.BaseURL()
	if baseURL == "" {
		return nil, nil, errors.New("尚未设置扬帆自动化测试平台域名")
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := baseURL + path

	var bodyReader io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if token := c.store.Token(); token != "" {
		req.Header.Set("x-token", token)
	}
	for k, v := range extraHeaders {
		if strings.TrimSpace(k) == "" {
			continue
		}
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.Header, nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.Header, respBody, errors.New(strings.TrimSpace(string(respBody)))
	}
	return resp.Header, respBody, nil
}
