package auth

import (
	"errors"
	"net/http"
	"strconv"

	"yangfan-ui/internal/config"
)

type Manager struct {
	store *config.Store
}

func New(store *config.Store) *Manager {
	return &Manager{store: store}
}

func (m *Manager) Token() string {
	if m.store == nil {
		return ""
	}
	return m.store.Token()
}

func (m *Manager) Set(token string, expiresAt int64) error {
	if m.store == nil {
		return errors.New("config store 未初始化")
	}
	return m.store.SetToken(token, expiresAt)
}

func (m *Manager) Clear() error {
	if m.store == nil {
		return errors.New("config store 未初始化")
	}
	return m.store.ClearAuth()
}

func (m *Manager) ApplyNewTokenFromHeaders(h http.Header) error {
	if m.store == nil {
		return errors.New("config store 未初始化")
	}
	token := h.Get("new-token")
	if token == "" {
		return nil
	}
	expiresAt := int64(0)
	if v := h.Get("new-expires-at"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			expiresAt = n
		}
	}
	return m.store.SetToken(token, expiresAt)
}
