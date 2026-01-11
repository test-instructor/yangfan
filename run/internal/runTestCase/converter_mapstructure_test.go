package runTestCase

import (
	"encoding/json"
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestConvertConfigToTConfigWithMapstructure(t *testing.T) {
	config := &platform.RunConfig{
		Name:    "test-config",
		BaseUrl: "https://example.com",
		Verify:  true,
		Variables: datatypes.JSONMap{
			"var1": "val1",
		},
		Headers: datatypes.JSONMap{
			"User-Agent": "hrp-test",
		},
		Timeout: 30,
	}

	environs := map[string]string{
		"ENV_KEY": "ENV_VAL",
	}

	tConfig := ConvertConfigToTConfigWithMapstructure(config, environs)

	assert.Equal(t, "test-config", tConfig.Name)
	assert.Equal(t, "https://example.com", tConfig.BaseURL)
	assert.True(t, tConfig.Verify)
	assert.Equal(t, "val1", tConfig.Variables["var1"])
	assert.Equal(t, "ENV_VAL", tConfig.Variables["ENV_KEY"])
	assert.Equal(t, "hrp-test", tConfig.Headers["User-Agent"])
	assert.Equal(t, float32(30), tConfig.RequestTimeout)
}

func TestConvertAutoStepToIStepWithMapstructure(t *testing.T) {
	req := &automation.Request{
		Method: "GET",
		URL:    "/api/test",
		Params: datatypes.JSONMap{"p1": "v1"},
		Headers: datatypes.JSONMap{"h1": "v1"},
	}

	autoStep := &automation.AutoStep{
		StepConfig: automation.StepConfig{
			StepName: "step1",
			Variables: datatypes.JSONMap{"svar": "sval"},
            Loops: 5,
		},
		Request: req,
	}
    
    // SetupHooks is JSON
    hooks := []string{"hook1", "hook2"}
    hooksBytes, _ := json.Marshal(hooks)
    autoStep.SetupHooks = datatypes.JSON(hooksBytes)

	iStep := ConvertAutoStepToIStepWithMapstructure(autoStep)

	step, ok := iStep.(*hrp.StepRequestWithOptionalArgs)
	assert.True(t, ok)
	assert.NotNil(t, step.StepRequest)
	assert.Equal(t, "step1", step.StepRequest.StepConfig.StepName)
	assert.Equal(t, "sval", step.StepRequest.StepConfig.Variables["svar"])
	assert.Equal(t, hrp.HTTPMethod("GET"), step.StepRequest.Request.Method)
	assert.Equal(t, "/api/test", step.StepRequest.Request.URL)
    assert.Equal(t, 5, step.StepRequest.StepConfig.Loops)
    
    // Check hooks
    assert.Equal(t, 2, len(step.StepRequest.StepConfig.SetupHooks))
    assert.Equal(t, "hook1", step.StepRequest.StepConfig.SetupHooks[0])
}
