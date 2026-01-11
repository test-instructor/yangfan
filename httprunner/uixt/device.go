package uixt

import (
	"bytes"

	"github.com/test-instructor/yangfan/httprunner/uixt/option"
	"github.com/test-instructor/yangfan/httprunner/uixt/types"
)

// current implemeted device: IOSDevice, AndroidDevice, HarmonyDevice
type IDevice interface {
	UUID() string
	NewDriver() (driver IDriver, err error)

	IsHealthy() (bool, error)

	Setup() error
	Teardown() error

	Install(appPath string, opts ...option.InstallOption) error
	Uninstall(packageName string) error

	ListPackages() ([]string, error)

	GetPackageInfo(packageName string) (types.AppInfo, error)
	ScreenShot() (*bytes.Buffer, error)
	// TODO: remove?
	LogEnabled() bool
}
