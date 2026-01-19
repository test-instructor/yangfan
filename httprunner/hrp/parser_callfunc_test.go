package hrp

import (
	"fmt"
	"testing"
)

type hasFalseButCallablePlugin struct {
	resultByName map[string]interface{}
}

func (p *hasFalseButCallablePlugin) Type() string { return "test" }

func (p *hasFalseButCallablePlugin) Path() string { return "" }

func (p *hasFalseButCallablePlugin) Has(name string) bool { return false }

func (p *hasFalseButCallablePlugin) Call(name string, args ...interface{}) (interface{}, error) {
	if v, ok := p.resultByName[name]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("function %s is not found", name)
}

func (p *hasFalseButCallablePlugin) Quit() error { return nil }

func (p *hasFalseButCallablePlugin) StartHeartbeat() {}

func TestCallFunc_PluginHasFalseButCallWorks(t *testing.T) {
	parser := NewParser()
	parser.Plugin = &hasFalseButCallablePlugin{
		resultByName: map[string]interface{}{
			"py_func": "ok",
		},
	}

	got, err := parser.CallFunc("py_func")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got != "ok" {
		t.Fatalf("expected %v, got %v", "ok", got)
	}
}
