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
	status, headers, respBody, err := c.doRaw(ctx, baseURL, method, path, reqBody, extraHeaders, c.store.Token())
	if err != nil {
		if status != 0 && len(respBody) > 0 && (status < 200 || status >= 300) {
			return headers, respBody, errors.New(strings.TrimSpace(string(respBody)))
		}
		return headers, respBody, err
	}
	return headers, respBody, nil
}

func (c *Client) DoWithBaseURL(ctx context.Context, baseURL string, method string, path string, reqBody any, extraHeaders map[string]string) (int, http.Header, []byte, error) {
	if strings.TrimSpace(baseURL) == "" {
		return 0, nil, nil, errors.New("baseURL 不能为空")
	}
	return c.doRaw(ctx, baseURL, method, path, reqBody, extraHeaders, "")
}

func (c *Client) doRaw(ctx context.Context, baseURL string, method string, path string, reqBody any, extraHeaders map[string]string, token string) (int, http.Header, []byte, error) {
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: 30 * time.Second}
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := strings.TrimRight(baseURL, "/") + path

	var bodyReader io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return 0, nil, nil, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return 0, nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if strings.TrimSpace(token) != "" {
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
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, resp.Header, nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.StatusCode, resp.Header, respBody, errors.New("http status not ok")
	}
	return resp.StatusCode, resp.Header, respBody, nil
}
