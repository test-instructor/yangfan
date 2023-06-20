package ios

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/test-instructor/yangfan/parsing/hrp/pkg/gidevice"
	"github.com/test-instructor/yangfan/parsing/hrp/pkg/uixt"
)

var iosRootCmd = &cobra.Command{
	Use:   "ios",
	Short: "simple utils for ios device management",
}

func getDevice(udid string) (gidevice.Device, error) {
	devices, err := uixt.GetIOSDevices(udid)
	if err != nil {
		return nil, err
	}
	if len(devices) > 1 {
		return nil, fmt.Errorf("found multiple attached devices, please specify ios udid")
	}
	return devices[0], nil
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(iosRootCmd)
}
