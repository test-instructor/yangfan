package ios

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
)

var xctestCmd = &cobra.Command{
	Use:   "xctest",
	Short: "run xctest",
	RunE: func(cmd *cobra.Command, args []string) error {
		if bundleID == "" {
			return fmt.Errorf("bundleID is required")
		}
		device, err := getDevice(udid)
		if err != nil {
			return err
		}

		global.GVA_LOG.Info("run xctest", zap.String("bundleID", bundleID))
		out, cancel, err := device.XCTest(bundleID)
		if err != nil {
			return errors.Wrap(err, "run xctest failed")
		}

		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

		// print xctest running logs
		go func() {
			for s := range out {
				fmt.Print(s)
			}
			done <- os.Interrupt
		}()

		<-done
		cancel()

		return nil
	},
}

var bundleID string

func init() {
	xctestCmd.Flags().StringVarP(&udid, "udid", "u", "", "filter by device's udid")
	xctestCmd.Flags().StringVarP(&bundleID, "bundleID", "b", "", "specify ios bundleID")
	iosRootCmd.AddCommand(xctestCmd)
}
