package runTestCase

import (
	"encoding/json"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/httprunner/uixt/option"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
)

// appendUIConfigToTConfig 将 UI 设备配置添加到 TConfig
func appendUIConfigToTConfig(tConfig *hrp.TConfig, config *platform.RunConfig) {
	if config.AndroidDeviceOptions != nil {
		tConfig.SetAndroid(convertAndroidOption(config.AndroidDeviceOptions))
	}
	if config.IOSDeviceOptions != nil {
		tConfig.SetIOS(convertIOSOption(config.IOSDeviceOptions))
	}
	if config.HarmonyDeviceOptions != nil {
		tConfig.SetHarmony(convertHarmonyOption(config.HarmonyDeviceOptions))
	}
	if config.BrowserDeviceOptions != nil {
		tConfig.SetBrowser(convertBrowserOption(config.BrowserDeviceOptions))
	}
}

func convertAndroidOption(opts *platform.AndroidDeviceOptions) option.AndroidDeviceOption {
	return func(o *option.AndroidDeviceOptions) {
		if opts.SerialNumber != nil {
			o.SerialNumber = *opts.SerialNumber
		}
		if opts.LogOn != nil {
			o.LogOn = *opts.LogOn
		}
		if opts.IgnorePopup != nil {
			o.IgnorePopup = *opts.IgnorePopup
		}
		if opts.AdbServerHost != nil {
			o.AdbServerHost = *opts.AdbServerHost
		}
		if opts.AdbServerPort != nil {
			o.AdbServerPort = int(*opts.AdbServerPort)
		}
		if opts.UIA2 != nil {
			o.UIA2 = *opts.UIA2
		}
		if opts.UIA2IP != nil {
			o.UIA2IP = *opts.UIA2IP
		}
		if opts.UIA2Port != nil {
			o.UIA2Port = int(*opts.UIA2Port)
		}
	}
}

func convertIOSOption(opts *platform.IOSDeviceOptions) option.IOSDeviceOption {
	return func(o *option.IOSDeviceOptions) {
		if opts.UDID != nil {
			o.UDID = *opts.UDID
		}
		if opts.LogOn != nil {
			o.LogOn = *opts.LogOn
		}
		if opts.IgnorePopup != nil {
			o.IgnorePopup = *opts.IgnorePopup
		}
		if opts.WDAPort != nil {
			o.WDAPort = int(*opts.WDAPort)
		}
		if opts.WDAMjpegPort != nil {
			o.WDAMjpegPort = int(*opts.WDAMjpegPort)
		}
	}
}

func convertHarmonyOption(opts *platform.HarmonyDeviceOptions) option.HarmonyDeviceOption {
	return func(o *option.HarmonyDeviceOptions) {
		if opts.ConnectKey != nil {
			o.ConnectKey = *opts.ConnectKey
		}
		if opts.LogOn != nil {
			o.LogOn = *opts.LogOn
		}
		if opts.IgnorePopup != nil {
			o.IgnorePopup = *opts.IgnorePopup
		}
	}
}

func convertBrowserOption(opts *platform.BrowserDeviceOptions) option.BrowserDeviceOption {
	return func(o *option.BrowserDeviceOptions) {
		if opts.BrowserID != nil {
			o.BrowserID = *opts.BrowserID
		}
		if opts.LogOn != nil {
			o.LogOn = *opts.LogOn
		}
		if opts.IgnorePopup != nil {
			o.IgnorePopup = *opts.IgnorePopup
		}
	}
}

// convertMobileStepToHrpStep 将 UI Step 转换为 HRP Step
func convertMobileStepToHrpStep(mobileStep *automation.MobileStep, platform string, stepConfig hrp.StepConfig) hrp.IStep {
	if mobileStep == nil {
		return nil
	}

	stepMobile := &hrp.StepMobile{
		StepConfig: stepConfig,
	}

	mobileUI := &hrp.MobileUI{
		Actions: convertMobileActions(mobileStep.Actions),
	}

	switch platform {
	case "android":
		stepMobile.Android = mobileUI
	case "ios":
		stepMobile.IOS = mobileUI
	case "harmony":
		stepMobile.Harmony = mobileUI
	case "browser":
		stepMobile.Browser = mobileUI
	}

	return stepMobile
}

func convertMobileActions(actions []automation.MobileAction) []option.MobileAction {
	var hrpActions []option.MobileAction
	for _, a := range actions {
		action := option.MobileAction{
			Method: option.ActionName(a.Method),
			Params: a.Params,
		}

		if a.Options != nil {
			action.Options = convertActionOptions(a.Options)
		}

		hrpActions = append(hrpActions, action)
	}
	return hrpActions
}

func convertActionOptions(opts *automation.ActionOptions) *option.ActionOptions {
	if opts == nil {
		return nil
	}
	// 使用 json 序列化/反序列化进行转换，避免手动映射大量字段
	// 因为结构体字段定义几乎一致
	data, err := json.Marshal(opts)
	if err != nil {
		return nil
	}
	var hrpOpts option.ActionOptions
	if err := json.Unmarshal(data, &hrpOpts); err != nil {
		return nil
	}
	return &hrpOpts
}
