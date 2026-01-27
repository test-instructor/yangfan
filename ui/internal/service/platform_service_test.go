package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/config"
	"yangfan-ui/internal/platformapi"
	"yangfan-ui/internal/platformclient"
)

func TestPlatformService_LoginWritesNodeAndStoresToken(t *testing.T) {
	var gotNode string
	loginSpec, err := platformapi.GetSpec(platformapi.EndpointLogin)
	if err != nil {
		t.Fatalf("GetSpec login: %v", err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != loginSpec.Path {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if v, ok := body["node"].(string); ok {
			gotNode = v
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"code": 0,
			"data": map[string]any{
				"user": map[string]any{
					"userName":      "admin",
					"authorityId":   888,
					"projectId":     1,
					"authorities":   []any{},
					"projectList":   []any{},
					"originSetting": map[string]any{},
				},
				"token":     "t1",
				"expiresAt": 123,
			},
			"msg": "ok",
		})
	}))
	defer srv.Close()

	appName := "yangfan-ui-test-login"
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("UserConfigDir: %v", err)
	}
	_ = os.RemoveAll(filepath.Join(cfgDir, appName))
	store, err := config.New(appName)
	if err != nil {
		t.Fatalf("config.New: %v", err)
	}
	if err := store.SetBaseURL(srv.URL); err != nil {
		t.Fatalf("SetBaseURL: %v", err)
	}
	authManager := auth.New(store)
	svc := NewPlatformService(platformclient.New(store), authManager)

	user, err := svc.Login(context.Background(), "admin", "pwd", "", "", "ui-node")
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if gotNode != "ui-node" {
		t.Fatalf("node want ui-node, got %q", gotNode)
	}
	if store.Token() != "t1" {
		t.Fatalf("token want t1, got %q", store.Token())
	}
	if user["userName"] != "admin" {
		t.Fatalf("user want admin, got %v", user["userName"])
	}
	_ = os.RemoveAll(filepath.Join(cfgDir, appName))
}

func TestPlatformService_SetUserAuthorityUpdatesTokenFromHeaders(t *testing.T) {
	var gotTokenOnGet string
	setAuthSpec, err := platformapi.GetSpec(platformapi.EndpointSetAuth)
	if err != nil {
		t.Fatalf("GetSpec setAuth: %v", err)
	}
	userInfoSpec, err := platformapi.GetSpec(platformapi.EndpointUserInfo)
	if err != nil {
		t.Fatalf("GetSpec userInfo: %v", err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case setAuthSpec.Path:
			w.Header().Set("new-token", "t2")
			w.Header().Set("new-expires-at", "999")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"code": 0,
				"data": map[string]any{},
				"msg":  "ok",
			})
		case userInfoSpec.Path:
			gotTokenOnGet = r.Header.Get("x-token")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"code": 0,
				"data": map[string]any{
					"userInfo": map[string]any{
						"userName":      "admin",
						"authorityId":   999,
						"projectId":     2,
						"authorities":   []any{},
						"projectList":   []any{},
						"originSetting": map[string]any{},
					},
				},
				"msg": "ok",
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	appName := "yangfan-ui-test-switch"
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("UserConfigDir: %v", err)
	}
	_ = os.RemoveAll(filepath.Join(cfgDir, appName))
	store, err := config.New(appName)
	if err != nil {
		t.Fatalf("config.New: %v", err)
	}
	if err := store.SetBaseURL(srv.URL); err != nil {
		t.Fatalf("SetBaseURL: %v", err)
	}
	if err := store.SetToken("t1", 0); err != nil {
		t.Fatalf("SetToken: %v", err)
	}
	authManager := auth.New(store)
	svc := NewPlatformService(platformclient.New(store), authManager)

	user, err := svc.SetUserAuthority(context.Background(), 999, 2)
	if err != nil {
		t.Fatalf("SetUserAuthority: %v", err)
	}
	if store.Token() != "t2" {
		t.Fatalf("token want t2, got %q", store.Token())
	}
	if gotTokenOnGet != "t2" {
		t.Fatalf("x-token want t2, got %q", gotTokenOnGet)
	}
	if user["projectId"] != float64(2) {
		t.Fatalf("user projectId want 2, got %v", user["projectId"])
	}
	_ = os.RemoveAll(filepath.Join(cfgDir, appName))
}
