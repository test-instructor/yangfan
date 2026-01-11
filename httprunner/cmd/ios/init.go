package ios

import (
	"github.com/spf13/cobra"

	"github.com/test-instructor/yangfan/httprunner/uixt"
	"github.com/test-instructor/yangfan/httprunner/uixt/option"
)

var CmdIOSRoot = &cobra.Command{
	Use:   "ios",
	Short: "simple utils for ios device management",
}

func getDevice(udid string) (*uixt.IOSDevice, error) {
	device, err := uixt.NewIOSDevice(option.WithUDID(udid))
	if err != nil {
		return nil, err
	}
	return device, nil
}
