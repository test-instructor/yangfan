package runTestCase

import (
	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"github.com/mitchellh/mapstructure"
)

// ConvertConfigToTConfigWithMapstructure uses mapstructure to convert RunConfig to TConfig
func ConvertConfigToTConfigWithMapstructure(config *platform.RunConfig, environs map[string]string) *hrp.TConfig {
	if config == nil {
		return hrp.NewConfig(config.Name)
	}

	// Convert struct to map
	data := StructToMap(config)

	tConfig := &hrp.TConfig{}

	// Decode map to TConfig
	// We use "json" tag because hrp structs likely use json tags.
	// Our StructToMap uses "mapstructure" tags from Source which we aligned with hrp's expected keys.
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           tConfig,
	})
	if err == nil {
		decoder.Decode(data)
	}

	// Inject environment variables into Variables
	if environs != nil {
		if tConfig.Variables == nil {
			tConfig.Variables = make(map[string]interface{})
		}
		for k, v := range environs {
			// Explicit variables have precedence
			if _, exists := tConfig.Variables[k]; !exists {
				tConfig.Variables[k] = v
			}
		}
	}

	return tConfig
}

// ConvertRequestToHrpRequestWithMapstructure converts Request to hrp.Request using mapstructure
func ConvertRequestToHrpRequestWithMapstructure(req *automation.Request) *hrp.Request {
	if req == nil {
		return nil
	}

	data := StructToMap(req)
	hrpReq := &hrp.Request{}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           hrpReq,
	})
	decoder.Decode(data)

	return hrpReq
}

// ConvertAutoStepToIStepWithMapstructure converts AutoStep to hrp.IStep using mapstructure
func ConvertAutoStepToIStepWithMapstructure(autoStep *automation.AutoStep) hrp.IStep {
	if autoStep == nil {
		return nil
	}

	// Convert entire AutoStep to Map (flattens embedded StepConfig)
	data := StructToMap(autoStep)

	// 1. Decode StepConfig
	stepConfig := hrp.StepConfig{}
	decoderConfig, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           &stepConfig,
	})
	decoderConfig.Decode(data)

	// 2. Decode Request
	hrpReq := &hrp.Request{}
	if reqData, ok := data["request"]; ok && reqData != nil {
		decoderReq, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			TagName:          "json",
			WeaklyTypedInput: true,
			Result:           hrpReq,
		})
		decoderReq.Decode(reqData)
	} else {
		// If no request data in map (maybe nil), handle graceful failure or return nil
		// converter.go warns "AutoStep has no request"
		return nil
	}

	stepRequest := &hrp.StepRequest{
		StepConfig: stepConfig,
		Request:    hrpReq,
	}

	// Handle upload logic specific to hrp
	if hrpReq.Upload != nil && len(hrpReq.Upload) > 0 {
		if stepRequest.Request.Headers == nil {
			stepRequest.Request.Headers = make(map[string]string)
		}
		stepRequest.Request.Headers["Content-Type"] = "${multipart_content_type($m_encoder)}"
		stepRequest.Request.Body = "$m_encoder"
	}

	return &hrp.StepRequestWithOptionalArgs{
		StepRequest: stepRequest,
	}
}
