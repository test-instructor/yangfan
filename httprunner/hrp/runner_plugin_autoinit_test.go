package hrp

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lingcetech/funplugin"
)

type stubPlugin struct {
	path string
}

func (p *stubPlugin) Type() string { return "test" }

func (p *stubPlugin) Path() string { return p.path }

func (p *stubPlugin) Has(name string) bool { return false }

func (p *stubPlugin) Call(name string, args ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("function %s is not found", name)
}

func (p *stubPlugin) Quit() error { return nil }

func (p *stubPlugin) StartHeartbeat() {}

func TestNewCaseRunner_AutoInitPluginWhenDebugtalkExists(t *testing.T) {
	tmpDir := t.TempDir()

	debugtalkPath := filepath.Join(tmpDir, PluginPySourceFile)
	if err := os.WriteFile(debugtalkPath, []byte("def foo():\n    return 1\n"), 0o644); err != nil {
		t.Fatalf("write debugtalk.py failed: %v", err)
	}

	casePath := filepath.Join(tmpDir, "demo.yml")
	if err := os.WriteFile(casePath, []byte(""), 0o644); err != nil {
		t.Fatalf("write testcase file failed: %v", err)
	}

	cfg := NewConfig("demo")
	cfg.Path = casePath
	testcase := TestCase{Config: cfg}

	called := 0
	originInit := initPluginFunc
	initPluginFunc = func(path, venv string, logOn bool) (funplugin.IPlugin, error) {
		called++
		return &stubPlugin{path: debugtalkPath}, nil
	}
	t.Cleanup(func() {
		initPluginFunc = originInit
	})

	caseRunner, err := NewCaseRunner(testcase, NewRunner(nil))
	if err != nil {
		t.Fatalf("NewCaseRunner returned error: %v", err)
	}
	if called != 1 {
		t.Fatalf("expected initPluginFunc called once, got %d", called)
	}
	if caseRunner.parser.Plugin == nil {
		t.Fatalf("expected Parser.Plugin not nil")
	}
	if cfg.PluginSetting == nil {
		t.Fatalf("expected config.PluginSetting not nil")
	}
	if cfg.PluginSetting.Path != debugtalkPath {
		t.Fatalf("expected plugin path %s, got %s", debugtalkPath, cfg.PluginSetting.Path)
	}
	if cfg.PluginSetting.Type != "py" {
		t.Fatalf("expected plugin type py, got %s", cfg.PluginSetting.Type)
	}
	if len(cfg.PluginSetting.Content) == 0 {
		t.Fatalf("expected plugin content not empty")
	}
}
