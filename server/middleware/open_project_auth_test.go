package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	"gorm.io/gorm"
)

func TestOpenProjectAuth_QueryOK(t *testing.T) {
	db := setupTestDB(t)
	global.GVA_DB = db
	p := projectmgr.Project{UUID: "u1", Secret: "s1"}
	if err := db.Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/open/runner/run", OpenProjectAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/open/runner/run?projectId=1&uuid=u1&secret=s1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestOpenProjectAuth_BodyOK(t *testing.T) {
	db := setupTestDB(t)
	global.GVA_DB = db
	p := projectmgr.Project{UUID: "u2", Secret: "s2"}
	if err := db.Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/open/runner/run", OpenProjectAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	body := []byte(`{"projectId":1,"uuid":"u2","secret":"s2"}`)
	req := httptest.NewRequest(http.MethodPost, "/open/runner/run", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestOpenProjectAuth_AliasOK(t *testing.T) {
	db := setupTestDB(t)
	global.GVA_DB = db
	p := projectmgr.Project{UUID: "u3", Secret: "s3"}
	if err := db.Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/open/runner/run", OpenProjectAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/open/runner/run?projectId=1&ci_uuid=u3&ci_secret=s3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestOpenProjectAuth_MissingProjectId(t *testing.T) {
	db := setupTestDB(t)
	global.GVA_DB = db

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/open/runner/run", OpenProjectAuth(), func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	req := httptest.NewRequest(http.MethodGet, "/open/runner/run?uuid=u&secret=s", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v, body=%s", err, w.Body.String())
	}
	if resp.Code == 0 {
		t.Fatalf("expected error code, got 0: %s", w.Body.String())
	}
}

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&projectmgr.Project{}); err != nil {
		t.Fatal(err)
	}
	return db
}
