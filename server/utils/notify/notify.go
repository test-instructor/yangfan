package notify

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type EnqueuedPayload struct {
	Event        string `json:"event"`
	ProjectId    uint   `json:"projectId"`
	TaskID       string `json:"task_id"`
	ReportID     uint   `json:"report_id"`
	NodeName     string `json:"node_name"`
	SendTimeUnix int64  `json:"send_time"`
	CaseType     string `json:"case_type"`
	CaseID       uint   `json:"case_id"`
	EnvID        int    `json:"env_id"`
	ConfigID     int    `json:"config_id"`
	RunMode      string `json:"run_mode"`
}

func SendCallback(ctx context.Context, callbackURL string, payload EnqueuedPayload) error {
	_, err := postJSON(ctx, callbackURL, payload)
	return err
}

func SendWebhook(ctx context.Context, webhookType, webhookURL, webhookSecret string, payload EnqueuedPayload) error {
	switch strings.ToLower(webhookType) {
	case "custom":
		_, err := postJSON(ctx, webhookURL, payload)
		return err
	case "feishu":
		return postFeishu(ctx, webhookURL, webhookSecret, payload)
	case "wecom":
		return postWeCom(ctx, webhookURL, payload)
	case "dingtalk":
		return postDingTalk(ctx, webhookURL, webhookSecret, payload)
	default:
		return errors.New("unsupported webhook_type")
	}
}

func postJSON(ctx context.Context, targetURL string, body any) ([]byte, error) {
	u, err := normalizeURL(targetURL)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, errors.New("webhook http status " + strconv.Itoa(resp.StatusCode))
	}
	return respBody, nil
}

func normalizeURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("invalid url scheme")
	}
	host := u.Hostname()
	if host == "" {
		return "", errors.New("invalid url host")
	}
	if ip := net.ParseIP(host); ip != nil {
		if ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
			return "", errors.New("invalid url host")
		}
	}
	return u.String(), nil
}

func formatEnqueuedText(p EnqueuedPayload) string {
	var parts []string
	parts = append(parts, "CI任务已入队")
	parts = append(parts, "projectId="+strconv.FormatUint(uint64(p.ProjectId), 10))
	parts = append(parts, "task_id="+p.TaskID)
	parts = append(parts, "report_id="+strconv.FormatUint(uint64(p.ReportID), 10))
	if p.NodeName != "" {
		parts = append(parts, "node="+p.NodeName)
	}
	return strings.Join(parts, " ")
}

func postFeishu(ctx context.Context, webhookURL, secret string, payload EnqueuedPayload) error {
	body := map[string]any{
		"msg_type": "text",
		"content": map[string]any{
			"text": formatEnqueuedText(payload),
		},
	}
	if secret != "" {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		body["timestamp"] = timestamp
		body["sign"] = feishuSign(timestamp, secret)
	}
	_, err := postJSON(ctx, webhookURL, body)
	return err
}

func feishuSign(timestamp, secret string) string {
	toSign := timestamp + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func postWeCom(ctx context.Context, webhookURL string, payload EnqueuedPayload) error {
	body := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": formatEnqueuedText(payload),
		},
	}
	_, err := postJSON(ctx, webhookURL, body)
	return err
}

func postDingTalk(ctx context.Context, webhookURL, secret string, payload EnqueuedPayload) error {
	targetURL := webhookURL
	if secret != "" {
		u, err := url.Parse(webhookURL)
		if err != nil {
			return err
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.New("invalid url scheme")
		}
		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		sign := dingTalkSign(timestamp, secret)
		q := u.Query()
		q.Set("timestamp", timestamp)
		q.Set("sign", sign)
		u.RawQuery = q.Encode()
		targetURL = u.String()
	}
	body := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": formatEnqueuedText(payload),
		},
	}
	_, err := postJSON(ctx, targetURL, body)
	return err
}

func dingTalkSign(timestamp, secret string) string {
	toSign := timestamp + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toSign))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}

func PostJSON(ctx context.Context, targetURL string, body any) ([]byte, error) {
	return postJSON(ctx, targetURL, body)
}

func NormalizeURL(raw string) (string, error) {
	return normalizeURL(raw)
}
